package labels

import (
	"os"
)

const (
	experimentLabelKey = "EXPERIMENT_LABEL"
	assignmentLabelKey = "ASSIGNMENT_LABEL"
	variantLabelKey    = "VARIANT_LABEL"

	// ExperimentLabelDefault is The default label name for label/tagging an experiment identifier
	// onto a resources
	ExperimentLabelDefault = "experiment"

	// AssignmentLabelDefault is the default label name for label/tagging an experiment identifier
	// onto a resources
	AssignmentLabelDefault = "assignment"

	// VariantLabelDefault is the default label name for label/tagging an experiment identifier
	// onto a resources
	VariantLabelDefault = "variant"
)

// Labels is a collection of labels
type Labels []Label

// Label is a k/v pair of label name/values
type Label struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// NewLabels returns a set of labels with names set by their defaults
func NewLabels() []Label {
	return Labels{
		Label{
			Name: ExperimentLabelName(),
		},
		Label{
			Name: AssignmentLabelName(),
		},
		Label{
			Name: VariantLabelName(),
		},
	}
}

// ExperimentLabelName returns the name value for the experiment label,
// and respects override configuration form environment variable
func ExperimentLabelName() string {
	if label := os.Getenv(experimentLabelKey); label != "" {
		return label
	}

	return ExperimentLabelDefault
}

// AssignmentLabelName returns the name value for the experiment label,
// and respects override configuration form environment variable
func AssignmentLabelName() string {
	if label := os.Getenv(assignmentLabelKey); label != "" {
		return label
	}

	return AssignmentLabelDefault
}

// VariantLabelName returns the name value for the experiment label,
// and respects override configuration from environment variable
func VariantLabelName() string {
	if label := os.Getenv(variantLabelKey); label != "" {
		return label
	}

	return VariantLabelDefault
}

func Compare() bool { return false }

func (l *Labels) Contains(label Label) bool {
	return false
}

func (l *Labels) Any() bool {
	return len(*l) > 0
}

func (l *Labels) Intersect(compare Labels) Labels {
	return Labels{}
}

func (l *Label) Diff(compare Labels) Labels {
	return Labels{}
}

func (l *Labels) Union(join Labels) Labels {
	return Labels{}
}
