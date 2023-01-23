package elasticache

import (
	"fmt"

	"github.com/edobtc/cloudkit/labels"
	"github.com/edobtc/cloudkit/resources/providers"
	"github.com/edobtc/cloudkit/target"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type modified []string

// Config holds allowed values
// for an implemented resource provider. Any value outside of this config
// is unable to be modified during an experiment
type Config struct {
	// InstanceType is the cluster compute resource
	InstanceType string

	// ClusterSize is the size of the cluster
	ClusterSize int64
}

// Provisioner implements an elasticache provisioner
type Provisioner struct {
	// Config holds our internal configuration options
	// for the instance of the provisioner
	Config Config

	// RemoteConfig identifies the
	RemoteConfig *elasticache.CacheCluster
}

// NewProvisioner initializes a provisioner
// with defaults
func NewProvisioner(yml []byte) providers.Provider {
	cfg := Config{}
	err := yaml.Unmarshal(yml, &cfg)

	if err != nil {
		return nil
	}

	return &Provisioner{Config: cfg}
}

// Select is similar to Read yet copies a selection of resources based on the Target configuration
func (p *Provisioner) Select() (target.Selection, error) { return target.Selection{}, nil }

// Read fetches and stores the configuration for an existing
// elasticache cluster. What is read of the existing resource acts
// as the template/configuration to implement a clone via creating a
// new resource with the existing output as input for a variant
func (p *Provisioner) Read() error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))

	svc := elasticache.New(sess)

	input := &elasticache.DescribeCacheClustersInput{
		CacheClusterId: aws.String("hieroglyph-e2e"),
	}

	result, err := svc.DescribeCacheClusters(input)

	if err != nil {
		log.Error(err)
	}

	p.RemoteConfig = result.CacheClusters[0]

	return nil
}

func (p *Provisioner) modify() (modified, error) {
	// iterate through supplied config
	// modify outgoing request
	// return a list of modified/dirtied fields

	return modified{}, nil
}

// Clone creates a modified variant
func (p *Provisioner) Clone() error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))

	svc := elasticache.New(sess)

	variantName := fmt.Sprintf("var-a-%s", *p.RemoteConfig.CacheClusterId)

	// This is pretty basic, but a working minimum to clone.
	// We'll have to build some smarts into detecting more of the existing config
	// (ie: vpc, subnets, etc) but it's not a great deal much more work to code this way.
	// we can also do a serialize -> deserialize from output to input format and ONLY
	// override the configurable params (ie: cluster size and instace type) and reduce this
	// all to JUST the serialization code.
	input := &elasticache.CreateCacheClusterInput{
		CacheClusterId: aws.String(variantName),
		CacheNodeType:  aws.String(p.Config.InstanceType),
		Engine:         p.RemoteConfig.Engine,
		EngineVersion:  p.RemoteConfig.EngineVersion,
		NumCacheNodes:  aws.Int64(p.Config.ClusterSize),
	}

	result, err := svc.CreateCacheCluster(input)

	if err != nil {
		log.Error(err)
		return err
	}

	log.Info(result)

	return nil
}

// ProbeReadiness checks that the provisioned resource is available and
// ready to be included in a live experiment
func (p *Provisioner) ProbeReadiness() (bool, error) {
	return false, nil
}

// Teardown eradicates any resource that has been
// provisioned as part of a variant
func (p *Provisioner) Teardown() error {
	// Needs to look up variants based on
	// labels / tags which identify a variant name, experiment,
	// and ideally a namespace
	return nil
}

// Apply runs the provisioner end to end, so calls
// read and clone
func (p *Provisioner) Apply() error { return nil }

// Cancel will abort and running or submitted provisioner
func (p *Provisioner) Cancel() error { return nil }

// Stop will stop any running provisioner
func (p *Provisioner) Stop() error { return nil }

// AwaitReadiness should be implemented to detect
// when a provisioner has finished setting up a variant
// and can begin using it in an experiment
func (p *Provisioner) AwaitReadiness() chan error { return make(chan error) }

// Annotate should implement applying labels or tags for a given resource type
func (p *Provisioner) Annotate(id string, l labels.Labels) error { return nil }
