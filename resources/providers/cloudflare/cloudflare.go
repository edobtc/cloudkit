package cloudflare

import (
	"github.com/edobtc/cloudkit/labels"
	"github.com/edobtc/cloudkit/resources/providers"
	"github.com/edobtc/cloudkit/target"
	"gopkg.in/yaml.v2"
)

// EXAMPLE##BOILERPLATE

// Config holds allowed values
// for an implemented resource provider. Any value outside of this config
// is unable to be modified during an experiment
type Config struct {
	// InstanceType is the cluster compute resource
	InstanceType string `yaml:"instanceType"`

	// ClusterSize is the size of the cluster
	ClusterSize int64 `yaml:"clusterSize"`
}

// CloudflareProvisioner implements a CloudflareProvisioner
type CloudflareProvisioner struct {
	// Config holds our internal configuration options
	// for the instance of the CloudflareProvisioner
	Config Config
}

// NewProvisioner initializes a CloudflareProvisioner
// with defaults
func NewProvisioner(yml []byte) providers.Provider {
	cfg := Config{}
	err := yaml.Unmarshal(yml, &cfg)

	if err != nil {
		return nil
	}

	return &CloudflareProvisioner{Config: cfg}
}

// Select is similar to Read yet copies a selection of resources based on the Target configuration
func (p *CloudflareProvisioner) Select() (target.Selection, error) { return target.Selection{}, nil }

// Read fetches and stores the configuration for an existing
// elasticache cluster. What is read of the existing resource acts
// as the template/configuration to implement a clone via creating a
// new resource with the existing output as input for a variant
func (p *CloudflareProvisioner) Read() error {
	return nil
}

// Clone creates a modified variant
func (p *CloudflareProvisioner) Clone() error {
	return nil
}

// ProbeReadiness checks that the provisioned resource is available and
// ready to be included in a live experiment
func (p *CloudflareProvisioner) ProbeReadiness() (bool, error) {
	return false, nil
}

// Teardown eradicates any resource that has been
// provisioned as part of a variant
func (p *CloudflareProvisioner) Teardown() error {
	// Needs to look up variants based on
	// labels / tags which identify a variant name, experiment,
	// and ideally a namespace
	return nil
}

// Apply runs the CloudflareProvisioner end to end, so calls
// read and clone
func (p *CloudflareProvisioner) Apply() error { return nil }

// Cancel will abort and running or submitted CloudflareProvisioner
func (p *CloudflareProvisioner) Cancel() error { return nil }

// Stop will stop any running CloudflareProvisioner
func (p *CloudflareProvisioner) Stop() error { return nil }

// AwaitReadiness should be implemented to detect
// when a CloudflareProvisioner has finished setting up a variant
// and can begin using it in an experiment
func (p *CloudflareProvisioner) AwaitReadiness() chan error { return make(chan error) }

// Annotate should implement applying labels or tags for a given resource type
func (p *CloudflareProvisioner) Annotate(id string, l labels.Labels) error { return nil }
