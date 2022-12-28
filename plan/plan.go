package plan

import (
	"github.com/edobtc/cloudkit/labels"
	"github.com/edobtc/cloudkit/resources/declaration"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	// CurrentVersion is the current API version
	CurrentVersion = "v1alpha1"
)

// Definition wraps the entire plan, namely version info
// type/kind info, metadata and the experiment Spec itself
type Definition struct {
	// APIVersion is the current version of the plan format
	APIVersion string `yaml:"apiVersion"`

	// Kind is what type of this definition is
	Kind string `yaml:"kind"`

	// Metadata houses metadata about the plan
	Metadata metadata `yaml:"metadata"`

	// Spec contains the specification for the experiment plan
	Spec declaration.Declaration `yaml:"spec"`
}

type metadata struct {
	Name   string        `yaml:"name"`
	labels labels.Labels `yaml:"labels"`
}

// NewDefinition initializes a Plan that can be
// used as the basis for starting a plan configuration
func NewDefinition() Definition {
	return Definition{
		APIVersion: CurrentVersion,
		Kind:       "experiment",
		Metadata: metadata{
			Name:   "name",
			labels: labels.NewLabels(),
		},
		Spec: declaration.NewDeclaration(),
	}
}

// ParseYAML de-serializes an experiment plan from
// submitted bytes
func ParseYAML(data []byte) (*Definition, error) {
	d := NewDefinition()

	err := yaml.Unmarshal(data, &d)

	if err != nil {
		return &d, errors.Wrap(err, "plan.ParseYML")
	}

	return &d, nil
}

// ToYAML serializes a experiment plan to YAML
func (p *Definition) ToYAML() ([]byte, error) {
	out, err := yaml.Marshal(p)

	if err != nil {
		return []byte{}, errors.Wrap(err, "plan.toYAML")
	}

	return out, nil
}
