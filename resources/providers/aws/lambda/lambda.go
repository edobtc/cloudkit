package lambda

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/edobtc/cloudkit/labels"
	"github.com/edobtc/cloudkit/resources/providers"
	"github.com/edobtc/cloudkit/resources/providers/aws/auth"
	"github.com/edobtc/cloudkit/target"

	"github.com/google/uuid"

	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"

	log "github.com/sirupsen/logrus"
)

const (
	// DefaultOperation is the default mode with which to perform a clone
	DefaultOperation = "clone"

	// DefaultCanaryWeight shall be an equal, uniform split
	DefaultCanaryWeight = 0.5

	DefaultSize = 512
)

var (
	// ErrOperationNotDefined is returned if we attempt to clone a function and the
	// given operation in the config is neither clone or alias
	ErrOperationNotDefined = errors.New("operation not defined, expecting clone or alias")
)

// Config holds allowed values
// for an implemented resource provider.
// Any value outside of this config
// is unable to be modified during an experiment
type Config struct {
	// Operation allows you to chose between
	// clone or alias, default is to alias
	Operation string

	// Canary defaults to false
	Canary bool

	Name string

	// CanaryWeight, which should be set from the
	CanaryWeight float64

	// MemorySize is the lambda compute resources size
	MemorySize int64

	Timeout int64

	// Handler allows a different handler to be used
	Handler string

	// Version of the deployed function
	Version string

	// Function runtime
	Runtime string
}

// Provider implements a lambda Provider
type Provider struct {
	Target target.Target

	// Config holds our internal configuration options
	// for the instance of the Provider
	Config Config

	// RemoteConfig identifies the
	RemoteConfig *lambda.GetFunctionOutput

	// CurrentAliasArn is the ARN of an alias if we are operating a clone
	CurrentAliasArn string
}

// NewProvider initializes a Provider
// with defaults
func NewProvider(req *pb.CreateRequest) providers.Provider {

	cfg := Config{
		Name:       req.Config.Name,
		MemorySize: sizeMap(req.Config.Size),
		Version:    req.Config.Version,
	}

	return &Provider{Config: cfg}
}

func sizeMap(size string) int64 {
	value, err := strconv.Atoi(size)

	if err != nil {
		return DefaultSize
	}

	return int64(value)
}

// Read fetches and stores the configuration for an existing
// lambda cluster. What is read of the existing resource acts
// as the template/configuration to implement a clone via creating a
// new resource with the existing output as input for a variant
func (p *Provider) Read() error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))

	svc := lambda.New(sess)

	input := &lambda.GetFunctionInput{
		FunctionName: aws.String(p.Target.ID),

		// Qualifier can specify a different version (ie: deployed but not released)
		// Qualifier: aws.String("$LATEST"),
	}

	result, err := svc.GetFunction(input)

	if err != nil {
		return err
	}

	p.RemoteConfig = result

	return nil
}

// Clone creates a modified variant
func (p *Provider) Clone() error {
	switch p.Config.Operation {
	case "clone":
		return p.clone()
	case "alias":
		return p.aliasClone()
	default:
		return ErrOperationNotDefined
	}
}

// clone duplicates the target function into a COMPLETELY new function
// resource with it's own ARN
func (p *Provider) clone() error {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))

	svc := lambda.New(sess)

	uuid := uuid.New()

	variantName := fmt.Sprintf("%s-%s", uuid, *p.RemoteConfig.Configuration.FunctionName)

	code, err := download(*p.RemoteConfig.Code.Location)
	if err != nil {
		return err
	}

	input := &lambda.CreateFunctionInput{
		Code: &lambda.FunctionCode{
			ZipFile: code,
		},
		FunctionName: aws.String(variantName),
		Handler:      p.RemoteConfig.Configuration.Handler,
		MemorySize:   aws.Int64(p.Config.MemorySize),
		Publish:      aws.Bool(true),
		Role:         p.RemoteConfig.Configuration.Role,
		Runtime:      p.RemoteConfig.Configuration.Runtime,
		Timeout:      aws.Int64(p.Config.Timeout),
		VpcConfig:    &lambda.VpcConfig{},
	}

	log.Info("Creating lambda clone")

	result, err := svc.CreateFunction(input)

	if err != nil {
		return err
	}

	log.Info(result)

	return nil
}

// aliasClone uses built in versioning + alias with optional traffic shifting
// to create a new version of the current function as a variant and optionally
// split traffic if Canary is set to true
//
// This requires a few steps to execute:
// First we need to make an explicit version of the targetted function as $LATEST can't be used in an alias,
// this is the controlFunction and snapshots the current function config
// Then we must modify the source function to apply the new config
// we then create another named version for the function with THAT variant config
// Then if Canary is true, we configure an alias of the two, determine proper weights based on if variants have configured
// weights and assign that to the traffic split. Invocations of the function that target the alias arn THEN
// will be split between the control and variant
func (p *Provider) aliasClone() error {
	experiment := fmt.Sprintf("%v-hiero", time.Now().Unix())

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))
	svc := lambda.New(sess)

	input := &lambda.PublishVersionInput{
		Description:  aws.String(fmt.Sprintf("%s.control", experiment)),
		FunctionName: aws.String(p.Target.ID),
	}

	controlFunction, err := svc.PublishVersion(input)

	if err != nil {
		return err
	}

	updateInput := &lambda.UpdateFunctionConfigurationInput{
		FunctionName: aws.String(p.Target.ID),
		MemorySize:   aws.Int64(p.Config.MemorySize),
	}

	_, err = svc.UpdateFunctionConfiguration(updateInput)
	if err != nil {
		return err
	}

	versionInput := &lambda.PublishVersionInput{
		Description:  aws.String(fmt.Sprintf("%s.variant", experiment)),
		FunctionName: aws.String(p.Target.ID),
	}

	variantFunction, err := svc.PublishVersion(versionInput)

	if err != nil {
		return err
	}

	if p.Config.Canary {
		if p.Config.CanaryWeight == 0.0 {
			p.Config.CanaryWeight = DefaultCanaryWeight
		}

		aliasInput := &lambda.CreateAliasInput{
			Name:            aws.String(experiment),
			Description:     aws.String(fmt.Sprintf("request group for hieroglyph %s", experiment)),
			FunctionName:    aws.String(p.Target.ID),
			FunctionVersion: controlFunction.Version,
			RoutingConfig: &lambda.AliasRoutingConfiguration{
				AdditionalVersionWeights: RoutingConfigurationWeights{
					fmt.Sprintf("%v", *variantFunction.Version): aws.Float64(p.Config.CanaryWeight), // get proper weighting and routes here
				},
			},
		}

		aliasConfig, err := svc.CreateAlias(aliasInput)

		if err != nil {
			return err
		}

		p.CurrentAliasArn = *aliasConfig.AliasArn
	}

	return nil
}

// ProbeReadiness checks that the provisioned resource is available and
// ready to be included in a live experiment
func (p *Provider) ProbeReadiness() (bool, error) {
	return false, nil
}

// Teardown eradicates any resource that has been
// provisioned as part of a variant
func (p *Provider) Teardown() error {
	// Needs to look up variants based on
	// labels / tags which identify a variant name, experiment,
	// and ideally a namespace
	return nil
}

func download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)

}

// Apply runs the Provider end to end, so calls
// read and clone
func (p *Provider) Apply() error {
	err := p.Read()
	if err != nil {
		return err
	}

	err = p.Clone()
	if err != nil {
		return err
	}

	return nil
}

// Cancel will abort and running or submitted Provider
func (p *Provider) Cancel() error { return nil }

// Stop will stop any running Provider
func (p *Provider) Stop() error { return nil }

// AwaitReadiness should be implemented to detect
// when a Provider has finished setting up a variant
// and can begin using it in an experiment
func (p *Provider) AwaitReadiness() chan error { return make(chan error) }

// Select is similar to Read yet copies a selection of resources based on the Target configuration
func (p *Provider) Select() (target.Selection, error) {
	selection := target.Selection{}

	sess, err := auth.Session()
	if err != nil {
		return selection, err
	}

	svc := lambda.New(sess)

	input := lambda.ListFunctionsInput{
		FunctionVersion: aws.String("ALL"),
	}

	results, err := svc.ListFunctions(&input)

	for _, function := range results.Functions {

		if *function.FunctionArn == p.Target.ID {
			selection.Add(target.Resource{
				ID:     *function.FunctionArn,
				Config: function,
			})
		} else {
			tagInput := lambda.ListTagsInput{
				Resource: function.FunctionArn,
			}

			if tags, err := svc.ListTags(&tagInput); err == nil {
				var matches bool

				if len(tags.Tags) > 0 && p.Target.Selectors.Any() {
					for key, value := range tags.Tags {
						matches = p.Target.Selectors.Contains(labels.Label{
							Name:  key,
							Value: *value,
						})
					}

					if matches {
						selection.Add(target.Resource{
							ID:     *function.FunctionArn,
							Config: function,
						})
					}
				}
			}
		}
	}

	return target.Selection{}, nil
}

// Annotate should implement applying labels or tags for a given resource type
func (p *Provider) Annotate(id string, l labels.Labels) error {
	sess, err := auth.Session()
	if err != nil {
		return err
	}

	svc := lambda.New(sess)

	tags := map[string]*string{}

	for _, selector := range l {
		tags[selector.Name] = aws.String(selector.Value)
	}

	input := &lambda.TagResourceInput{
		Resource: aws.String(id),
		Tags:     tags,
	}

	result, err := svc.TagResource(input)
	if err != nil {
		return err
	}

	log.Info(result)

	return nil
}

// RoutingConfigurationWeights is for configuration of routing weights if
// Canary is true
type RoutingConfigurationWeights map[string]*float64
