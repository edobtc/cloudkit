package ec2

import (
	"fmt"

	"github.com/edobtc/cloudkit/labels"
	"github.com/edobtc/cloudkit/resources/providers"
	"github.com/edobtc/cloudkit/resources/providers/aws/auth"
	"github.com/edobtc/cloudkit/target"
	"gopkg.in/yaml.v2"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"

	log "github.com/sirupsen/logrus"
)

type modified []string

// Config holds allowed values for an implemented
// resource provider. Any value outside of this config
// is unable to be modified during an experiment
type Config struct {
	Target target.Target `yaml:"target"`

	// AMI is the amazon machine image type to use
	AMI string `yaml:"ami"`

	// InstanceType is the AWS compute instance type to use
	InstanceType string `yaml:"instanceType"`

	// Count is the size of the cluster
	Count int64 `yaml:"count"`
}

// Provisioner implements an ec2 provisioner
type Provisioner struct {
	// Config holds our internal configuration options
	// for the instance of the provisioner
	Config Config

	// RemoteConfig identifies the remote config
	RemoteConfig string
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

// Read fetches and stores the configuration for an existing
// ec2 instance. What is read of the existing resource acts
// as the template/configuration to implement a clone via creating a
// new resource with the existing output as input for a variant
func (p *Provisioner) Read() error {
	sess, err := auth.Session()
	if err != nil {
		return err
	}

	svc := ec2.New(sess)

	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(p.Config.Target.ID),
		},
	}
	result, err := svc.DescribeInstances(input)

	if err != nil {
		log.Error(err)
	}

	p.RemoteConfig = result.GoString()

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
	sess, err := auth.Session()
	if err != nil {
		return err
	}

	svc := ec2.New(sess)

	variantName := fmt.Sprintf("var-a-%s", p.Config.AMI)

	input := &ec2.RunInstancesInput{
		ImageId:      aws.String(p.Config.AMI),
		InstanceType: aws.String(p.Config.InstanceType),
		MinCount:     aws.Int64(p.Config.Count),
		MaxCount:     aws.Int64(p.Config.Count),

		// TODO: Need to ensure we apply and intersect experiment tags
		// here too
		TagSpecifications: []*ec2.TagSpecification{
			&ec2.TagSpecification{
				ResourceType: aws.String("instance"),
				Tags: []*ec2.Tag{
					&ec2.Tag{
						Key:   aws.String("name"),
						Value: aws.String(variantName),
					},
				},
			},
		},
	}

	result, err := svc.RunInstances(input)

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
// read and clone, or Annotate
func (p *Provisioner) Apply() error {
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

// Cancel will abort and running or submitted provisioner
func (p *Provisioner) Cancel() error { return nil }

// Stop will stop any running provisioner
func (p *Provisioner) Stop() error { return nil }

// AwaitReadiness should be implemented to detect
// when a provisioner has finished setting up a variant
// and can begin using it in an experiment
func (p *Provisioner) AwaitReadiness() chan error {
	ready := make(chan error)

	return ready
}

// Select is similar to Read yet copies a selection of resources based on
// the Target configuration
func (p *Provisioner) Select() (target.Selection, error) {
	selection := target.Selection{}

	sess, err := auth.Session()
	if err != nil {
		return selection, err
	}

	svc := ec2.New(sess)

	result, err := svc.DescribeInstances(p.filters())

	if err != nil {
		return selection, err
	}

	for _, r := range result.Reservations {
		for _, instance := range r.Instances {
			selection.Add(target.Resource{
				ID:     *instance.InstanceId,
				Config: instance,
			})
		}
	}

	return selection, nil
}

func (p *Provisioner) filters() *ec2.DescribeInstancesInput {
	filters := []*ec2.Filter{}

	if p.Config.Target.ID != "" {
		filters = append(filters, &ec2.Filter{
			Name: aws.String("instance-id"),
			Values: []*string{
				aws.String(p.Config.Target.ID),
			},
		})
	}

	if p.Config.Target.Selectors.Any() {
		for _, selector := range p.Config.Target.Selectors.ToLabels() {
			name := fmt.Sprintf("tag:%s", selector.Name)

			filters = append(filters, &ec2.Filter{
				Name: aws.String(name),
				Values: []*string{
					aws.String(selector.Value),
				},
			})
		}
	}

	return &ec2.DescribeInstancesInput{Filters: filters}
}

// Annotate should implement applying labels or tags for a given resource type
func (p *Provisioner) Annotate(id string, l labels.Labels) error {
	sess, err := auth.Session()
	if err != nil {
		return err
	}

	svc := ec2.New(sess)

	tags := []*ec2.Tag{}

	for _, selector := range l {
		tags = append(tags, &ec2.Tag{
			Key:   aws.String(selector.Name),
			Value: aws.String(selector.Value),
		})
	}

	input := &ec2.CreateTagsInput{
		Resources: []*string{
			aws.String(id),
		},
		Tags: tags,
	}

	result, err := svc.CreateTags(input)
	if err != nil {
		return err
	}

	log.Info(result)

	return nil
}
