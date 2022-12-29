package providers

import (
	"github.com/edobtc/cloudkit/labels"
	"github.com/edobtc/cloudkit/target"
)

type GenericConfig map[string]interface{}

// Provider is an interface which any specific
// integration should implement
type Provider interface {
	// Select will collect a list of resources based on a target configuration
	Select() (target.Selection, error)

	// Read will read the settings of a resource
	// by its key identifiers
	Read() error

	// Clone will copy an active resource
	// replacing any parameters with any that are present and
	// altered in the Provider's supplied config
	Clone() error

	// Apply will provision a resource from some declaration
	// or config that has been passed to the Provider
	Apply() error

	// Annotate applies Labels or Tags for a given resource
	Annotate(id string, l labels.Labels) error

	// ProbeReadiness is for testing that a
	// resource is ready to be enabled
	ProbeReadiness() (bool, error)

	// AwaitReadiness utilizes
	AwaitReadiness() chan error

	// Teardown is the de-provisioning
	// any resource created by the provider
	Teardown() error

	// Cancel will safely cancel any
	// in flight provision request as cleanly
	// as possible
	Cancel() error

	// Stop...i dunno yet
	Stop() error

	// start
	// restart
	// connect
}

// EXAMPLE##BOILERPLATE

// // Config holds allowed values
// // for an implemented resource provider. Any value outside of this config
// // is unable to be modified during an experiment
// type Config struct {
// 	// InstanceType is the cluster compute resource
// 	InstanceType string `yaml:"instanceType"`

// 	// ClusterSize is the size of the cluster
// 	ClusterSize int64 `yaml:"clusterSize"`
// }

// // Provider implements a Provider
// type Provider struct {
// 	// Config holds our internal configuration options
// 	// for the instance of the Provider
// 	Config Config
// }

// // NewProvider initializes a Provider
// // with defaults
// func NewProvider(yml []byte) providers.Provider {
// 	cfg := Config{}
// 	err := yaml.Unmarshal(yml, &cfg)

// 	if err != nil {
// 		return nil
// 	}

// 	return &Provider{Config: cfg}
// }

// // Select is similar to Read yet copies a selection of resources based on the Target configuration
// func (p *Provider) Select() (target.Selection, error) { return target.Selection{}, nil }

// // Read fetches and stores the configuration for an existing
// // elasticache cluster. What is read of the existing resource acts
// // as the template/configuration to implement a clone via creating a
// // new resource with the existing output as input for a variant
// func (p *Provider) Read() error {
// 	return nil
// }

// // Clone creates a modified variant
// func (p *Provider) Clone() error {
// 	return nil
// }

// // ProbeReadiness checks that the provisioned resource is available and
// // ready to be included in a live experiment
// func (p *Provider) ProbeReadiness() (bool, error) {
// 	return false, nil
// }

// // Teardown eradicates any resource that has been
// // provisioned as part of a variant
// func (p *Provider) Teardown() error {
// 	// Needs to look up variants based on
// 	// labels / tags which identify a variant name, experiment,
// 	// and ideally a namespace
// 	return nil
// }

// // Apply runs the Provider end to end, so calls
// // read and clone
// func (p *Provider) Apply() error { return nil }

// // Cancel will abort and running or submitted Provider
// func (p *Provider) Cancel() error { return nil }

// // Stop will stop any running Provider
// func (p *Provider) Stop() error { return nil }

// // AwaitReadiness should be implemented to detect
// // when a Provider has finished setting up a variant
// // and can begin using it in an experiment
// func (p *Provider) AwaitReadiness() chan error { return make(chan error) }

// // Annotate should implement applying labels or tags for a given resource type
// func (p *Provider) Annotate(id string, l labels.Labels) error { return nil }
