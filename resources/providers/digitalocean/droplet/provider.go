package droplet

import (
	"github.com/edobtc/cloudkit/labels"
	"github.com/edobtc/cloudkit/resources/providers"
	"github.com/edobtc/cloudkit/target"

	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"

	log "github.com/sirupsen/logrus"
)

// Config holds allowed values
// for an implemented resource provider. Any value outside of this config
// is unable to be modified during an experiment
type Config struct {
	ID      int    `yaml:"id"`
	Name    string `yaml:"name"`
	Alias   string `yaml:"alias"`
	Size    string `yaml:"size"`
	SSHKey  string `yaml:"sshKey"`
	VPC     string `yaml:"vpc"`
	ImageID string `yaml:"_"`

	// Configurations
	LND LND `yaml:"lnd"`
}

type LND struct {
	ImageID string
	Config  string `yaml:"config"`
}

// Provider implements a Provider
type Provider struct {
	// Config holds our internal configuration options
	// for the instance of the Provider
	Config Config
}

// NewProvider initializes a Provider
// with defaults
func NewProvider(req *pb.CreateRequest) providers.Provider {
	// maybe change these mappings eventually
	cfg := Config{
		Name:    req.Config.Name,
		Alias:   req.Config.Name,
		Size:    req.Config.Size,
		ImageID: req.Config.Version,
	}

	return &Provider{Config: cfg}
}

// Select is similar to Read yet copies a selection of resources based on the Target configuration
func (p *Provider) Select() (target.Selection, error) { return target.Selection{}, nil }

// Read fetches and stores the configuration for an existing
// elasticache cluster. What is read of the existing resource acts
// as the template/configuration to implement a clone via creating a
// new resource with the existing output as input for a variant
func (p *Provider) Read() error {
	log.Info("read")
	return nil
}

// Clone creates a modified variant
func (p *Provider) Clone() error {
	log.Info("clone")
	return nil
}

// ProbeReadiness checks that the provisioned resource is available and
// ready to be included in a live experiment
func (p *Provider) ProbeReadiness() (bool, error) {
	log.Info("probe readiness")
	return false, nil
}

// Teardown eradicates any resource that has been
// provisioned as part of a variant
func (p *Provider) Teardown() error {
	// Needs to look up variants based on
	// labels / tags which identify a variant name, experiment,
	// and ideally a namespace
	log.Info("teardown")
	return nil
}

// Apply runs the Provider end to end, so calls
// read and clone
func (p *Provider) Apply() error {
	log.Info("apply")
	return nil
}

// Cancel will abort and running or submitted Provider
func (p *Provider) Cancel() error {
	log.Info("cancel")
	return nil
}

// Stop will stop any running Provider
func (p *Provider) Stop() error {
	log.Info("stop")
	return nil
}

// AwaitReadiness should be implemented to detect
// when a Provider has finished setting up a variant
// and can begin using it in an experiment
func (p *Provider) AwaitReadiness() chan error {
	log.Info("await readiness")
	return make(chan error)
}

// Annotate should implement applying labels or tags for a given resource type
func (p *Provider) Annotate(id string, l labels.Labels) error {
	log.Info("annotate")
	return nil
}
