package declaration

import (
	"fmt"

	"github.com/edobtc/cloudkit/assignment"
	"github.com/edobtc/cloudkit/crypto"
	"github.com/edobtc/cloudkit/guard"
	"github.com/edobtc/cloudkit/labels"
	"github.com/edobtc/cloudkit/parameters"
	"github.com/edobtc/cloudkit/probe"
	"github.com/edobtc/cloudkit/target"
	"github.com/edobtc/cloudkit/variant"

	"github.com/google/uuid"
)

const (
	// DefaultOrchestrator is The default orchestrator to
	// use to run the Declaration
	DefaultOrchestrator = "split"

	// DefaultAssignment The default assignment method to use
	// when classifying parameter sets into a variant group
	DefaultAssignment = "uniform"
)

// Declaration is the definition of a designed
// Declaration
type Declaration struct {
	ID uuid.UUID `json:"id"`

	Orchestrator string `yaml:"orchestrator" json:"orchestrator"`

	Target target.Target `yaml:"target" json:"target"`

	// Declaration Name
	Name string `yaml:"name" json:"name"`

	// Salt
	Salt string `yaml:"salt" json:"salt"`

	// Assignment Algorithm to use
	// Defaults to uniform
	Assignment string `yaml:"assignment" json:"assignment"`

	Operator Assignment `json:"-" yaml:"-"`

	// Event is a name for annotating
	// and labeling Declaration data that is
	// emitted and relevant to the running Declaration
	Event string `yaml:"event" json:"event"`

	// Control alerts that we want to run an A/A
	// defaults to false
	Control bool `yaml:"control" json:"control"`

	// Provider is the target implementation
	// to use (ie: ec2, lambda, k8s)
	Provider string `yaml:"provider" json:"provider"`

	// Label all resources with the Declaration and variant
	// configuration, defaults to true
	AutoLabel bool `yaml:"autolabel" json:"autolabel"`

	// Parameters are the input parameter fields to be used
	// for segmenting and bucketing users
	Parameters parameters.Parameters `yaml:"parameters" json:"parameters"`

	// Probes are a collection of probes to evaluate the state of
	// the Declaration
	Probes []*probe.Probe `yaml:"probes" json:"probes"`

	// Guards are a collection of guards with which to
	// protect the safety and control of the Declaration
	Guards []*guard.Guard `yaml:"guards" json:"guards"`

	Variants []*variant.Variant `yaml:"variants" json:"variants"`
}

// NewDeclaration returns a instantiated Declaration with defaults
func NewDeclaration() Declaration {
	return Declaration{
		ID:           uuid.New(),
		AutoLabel:    true,
		Orchestrator: DefaultOrchestrator,
		Assignment:   DefaultAssignment,
		Salt:         crypto.GenerateSalt(),
	}
}

// String returns the segments of the hashable
// string that represents input to the assignment operator
// for randomization
func (e *Declaration) String() string {
	return fmt.Sprintf("%s.%s", e.Name, e.Salt)
}

// Assignment is the interface any randomization
// operator should implement to be pluggable into
// an Declaration such that the operator type can be
// replaced
type Assignment interface {
	Assign(hash uint64) (*variant.Variant, error)
}

func (e *Declaration) loadOperator() Assignment {
	switch e.Assignment {
	case "weighted", "weighted_choice", "weightedChoice":
		return assignment.NewWeightedChoice(e.Variants)
	case "sample", "sample_choice", "sampleChoice":
		return assignment.NewSample(e.Variants)
	case "probabilistic", "probabilistic_choice", "probabilisticChoice":
		return assignment.NewProbabilistic(e.Variants)
	}

	return assignment.NewUniformChoice(e.Variants)
}

func (e *Declaration) segmentString(p *parameters.Parameters) string {
	e.Operator = e.loadOperator()

	DeclarationString := e.String()
	paramString := p.WhitelistedString(e.Parameters)
	return fmt.Sprintf("%s.%s", DeclarationString, paramString)
}

// HashString returns the hash string representation of the sha1 crypto
// value for the full list of segments for using to determine an
// assignment
func (e *Declaration) HashString(p parameters.Parameters) string {
	return crypto.HashString(e.segmentString(&p))
}

// Hash returns the hash int64 representation of the sha1 crypto
// value for the full list of segments for using to determine an
// assignment
func (e *Declaration) Hash(p parameters.Parameters) uint64 {
	return crypto.Hash(e.segmentString(&p))
}

// Assign takes a specific parameter set and returns the variant assignment for
// said parameter set
func (e *Declaration) Assign(p parameters.Parameters) (*variant.Variant, error) {
	hash := crypto.Hash(e.segmentString(&p))

	variant, err := e.Operator.Assign(hash)
	if err != nil {
		return nil, err
	}

	variant.Metadata.Labels = labels.Labels{
		labels.Label{
			Name:  "", // labels.DeclarationLabelName(),
			Value: e.String(),
		},
		labels.Label{
			Name: labels.AssignmentLabelName(),
			// TODO: need to do some work here to not call this twice when setting
			// metadata, fine for now while everything is in flux but we'll optimize this
			// soon - @ramin
			Value: e.HashString(p),
		},
		labels.Label{
			Name:  labels.VariantLabelName(),
			Value: variant.Name,
		},
	}

	return variant, nil
}
