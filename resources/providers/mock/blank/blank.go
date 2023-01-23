package blank

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
	Name string

	// Size is the cluster compute resource
	Size string

	Version string
}

// Provisioner implements a provisioner
type Provisioner struct {
	// Config holds our internal configuration options
	// for the instance of the provisioner
	Config Config
}

// NewProvisioner initializes a provisioner
// with defaults
func NewProvisioner(req *pb.CreateRequest) providers.Provider {
	cfg := Config{
		Name:    req.Config.Name,
		Size:    req.Config.Size,
		Version: req.Config.Version,
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
	log.Info("read")
	return nil
}

// Clone creates a modified variant
func (p *Provisioner) Clone() error {
	log.Info("clone")
	return nil
}

// ProbeReadiness checks that the provisioned resource is available and
// ready to be included in a live experiment
func (p *Provisioner) ProbeReadiness() (bool, error) {
	log.Info("probe readiness")
	return false, nil
}

// Teardown eradicates any resource that has been
// provisioned as part of a variant
func (p *Provisioner) Teardown() error {
	// Needs to look up variants based on
	// labels / tags which identify a variant name, experiment,
	// and ideally a namespace
	log.Info("teardown")
	return nil
}

// Apply runs the provisioner end to end, so calls
// read and clone
func (p *Provisioner) Apply() error {
	log.Info("apply")
	return nil
}

// Cancel will abort and running or submitted provisioner
func (p *Provisioner) Cancel() error {
	log.Info("cancel")
	return nil
}

// Stop will stop any running provisioner
func (p *Provisioner) Stop() error {
	log.Info("stop")
	return nil
}

// AwaitReadiness should be implemented to detect
// when a provisioner has finished setting up a variant
// and can begin using it in an experiment
func (p *Provisioner) AwaitReadiness() chan error {
	log.Info("await readiness")
	return make(chan error)
}

// Annotate should implement applying labels or tags for a given resource type
func (p *Provisioner) Annotate(id string, l labels.Labels) error {
	log.Info("annotate")
	return nil
}
