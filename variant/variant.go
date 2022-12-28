package variant

import (
	"github.com/edobtc/cloudkit/labels"
)

// Variant defines a choice/or variant option in an experiment
type Variant struct {
	//
	// Randomization Assignment Paramaters
	//
	// Weight, per variant, for weighted choice
	Weight float64 `yaml:"weight" json:"weight"`

	// P Value if we're doing probabalistic
	P *float32 `yaml:"p" json:"p"`

	// Drawer Count for Sample
	DrawCount *int `yaml:"drawCount" json:"drawCount"`

	// Metadata represents some metadata/labels we should
	// Apply to provisioned variants to identify them
	Metadata metadata `yaml:"metadata" json:"metadata"`

	// Config represents config for
	// the variant, which relies specifically
	// on the implemented configuration options
	// for a provider which can be targetted
	Config map[string]interface{} `yaml:"config" json:"config"`

	// Spec represents some configurable/mergable
	// set of configuration for merging with an existing
	// deployed spec for a resource such as a pod
	Spec map[string]interface{} `yaml:"spec" json:"spec"`

	// Label all resources with the experiment and variant
	// configuration, defaults to true
	AutoLabel bool `yaml:"autolabel" json:"autolabel"`

	//ExperimentId is the experiment the variant belongs to
	ExperimentID string

	// Id (temporary) is an internal id for addressing the variant
	ID string `yaml:"id" json:"id"`

	// AssignmentHash is set when a variant is returned as part of an
	// assignment operation
	AssignmentHash string `yaml:"assignmentHash" json:"assignmentHash"`

	// Name (optional) is a human readable format for the variant
	Name string `yaml:"name" json:"name"`
}

// NewVariant returns an initialized Variant
// with proper defaults
func NewVariant() Variant {
	return Variant{
		AutoLabel: true,
		Metadata: metadata{
			Labels: labels.Labels{},
		},
	}
}

type metadata struct {
	Labels labels.Labels
}
