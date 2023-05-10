package security_groups

import (
	"github.com/edobtc/cloudkit/labels"
	"github.com/edobtc/cloudkit/resources/providers"

	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"
)

// SecurityGroupsProvider implements a SecurityGroupsProvider
type SecurityGroupsProvider struct {
	// Config holds our internal configuration options
	// for the instance of the SecurityGroupsProvider
	Config *Config

	SGs []interface{}
}

// NewProvider initializes a SecurityGroupsProvider
// with defaults
func NewProvider(req *pb.CreateRequest) providers.Provider {
	return &SecurityGroupsProvider{
		Config: &Config{},
	}
}

func NewProviderFromConfig(cfg *Config) (providers.Provider, error) {
	return &SecurityGroupsProvider{
		Config: MergeDefaultConfig(cfg),
	}, nil
}

func (p *SecurityGroupsProvider) Apply() error {
	return nil
}

// Read fetches and stores the configuration for an existing
// SecurityGroups zone. What is read of the existing resource acts
// as the template/configuration to implement a clone via creating a
// new resource with the existing output as input for a variant
func (p *SecurityGroupsProvider) Read() error {
	return nil
}

// Clone creates a modified variant
func (p *SecurityGroupsProvider) Clone() error {
	return nil
}

// ProbeReadiness checks that the provisioned resource is available and
// ready to be included in a request
func (p *SecurityGroupsProvider) ProbeReadiness() (bool, error) {
	return false, nil
}

// Teardown eradicates any resource that has been
// provisioned as part of a variant
func (p *SecurityGroupsProvider) Teardown() error {
	// Needs to look up variants based on
	// labels / tags which identify a variant name, experiment,
	// and ideally a namespace
	return nil
}

// Cancel will abort and running or submitted SecurityGroupsProvider
func (p *SecurityGroupsProvider) Cancel() error { return nil }

// Stop will stop any running SecurityGroupsProvider
func (p *SecurityGroupsProvider) Stop() error { return nil }

// AwaitReadiness should be implemented to detect
// when a SecurityGroupsProvider has finished setting up a variant
// and can begin using it in an experiment
func (p *SecurityGroupsProvider) AwaitReadiness() chan error { return make(chan error) }

// Annotate should implement applying labels or tags for a given resource type
func (p *SecurityGroupsProvider) Annotate(id string, l labels.Labels) error { return nil }
