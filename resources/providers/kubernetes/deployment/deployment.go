package deployment

import (
	"github.com/edobtc/cloudkit/labels"
	"github.com/edobtc/cloudkit/resources/providers"
	"github.com/edobtc/cloudkit/target"
	"gopkg.in/yaml.v2"
)

type modified []string

// Config holds allowed values
// for an implemented resource provider. Any value outside of this config
// is unable to be modified during an experiment
type Config struct {
}

// Provisioner implements an k8s/deployment provisioner
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
// deployment instance. What is read of the existing resource acts
// as the template/configuration to implement a clone via creating a
// new resource with the existing output as input for a variant
func (p *Provisioner) Read() error {
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
	return nil
}

// ProbeReadiness checks that the provisioned resource is available and
// ready to be included in a live experiment
func (p *Provisioner) ProbeReadiness() (bool, error) {
	return false, nil
}

// Teardown eradicates any resource that“ has been
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

// Select is similar to Read yet copies a selection of resources based on the Target configuration
func (p *Provisioner) Select() (target.Selection, error) { return target.Selection{}, nil }

// Annotate should implement applying labels or tags for a given resource type
func (p *Provisioner) Annotate(id string, l labels.Labels) error { return nil }
