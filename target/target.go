package target

import (
	"github.com/edobtc/cloudkit/labels"
	"github.com/edobtc/cloudkit/parameters"
)

const (
	// DefaultNamespace is the default value when no namespace is set
	DefaultNamespace string = "default"
)

// Target is a set of rules for targetting a resource, or collection of resources
// for performing some experiment operation
type Target struct {
	// ID (optional) should be set if we
	// are targeting an existing resource (ie: lambda, ec2)
	// that we want to make a clone of for seeding the base configuration
	// of variants with
	//
	// +optional
	ID string `json:"id,omitempty"`

	// Namespace represents a specific namespace you want to target
	// resources in, usually only applicable to kubernetes
	//
	// +optional
	Namespace string

	// Resource (optional) a specific external resource
	// such as ec2, lambda
	//
	// +optional
	Resource string `json:"resource,omitempty"`

	// Kind (optional) represents a Kubernetes specific kind/resource
	//
	// +optional
	Kind string `json:"kind,omitempty"`

	// Selectors
	// +optional
	Selectors *selectors `yaml:"selectors,omitempty" json:"selectors,omitempty"`

	// Count is the number to select from a collection target
	Count int

	// Selection is the criteria for fulfilling the count number when
	// we are selecting a collection and that size exceeds the count
	// possible options are random, oldest, newest
	Selection string
}

// Empty returns true if all selector rules are unset or nil, ie:
// there are no selection criteria
func (t *Target) Empty() bool {
	return t == nil || (!t.hasSelectors() == false && t.ID == "")
}

func (t *Target) hasSelectors() bool {
	if t.Selectors == nil || t.Selectors.Any() == false {
		return false
	}

	return true
}

// Single returns true if we are targeting a single resource by its
// unique identifier
func (t *Target) Single() bool {
	return t.ID != ""
}

// SafeNamespace returns the set namespace or 'default'
// in the case it is not set
func (t *Target) SafeNamespace() string {
	if t.Namespace == "" {
		return DefaultNamespace
	}

	return t.Namespace
}

// DeepCopyInto is generated, don't edit
func (t *Target) DeepCopyInto(out *Target) {
	*out = *t
}

// DeepCopy is generated, don't edit
func (t *Target) DeepCopy() *Target {
	if t == nil {
		return nil
	}
	out := new(Target)
	t.DeepCopyInto(out)
	return out
}

// Resources are a collection of Resource items
type Resources []Resource

// Resource is a general representation of some kind of resource
// which a provider might handle (ec2 instance, k8s pod, lambda function)
type Resource struct {
	Name   string
	ID     string
	Meta   interface{}
	Labels labels.Labels
	Config interface{}
}

// Selection is the selection result of a select operation
// which uses target rules to grab a set of resources
type Selection struct {
	Data     Resources
	Selected Resources
}

// Select will take a list of resources in data, and apply selection criteria
// to those in order to segment and select a set of resources to include in selection
func (s *Selection) Select(p parameters.Parameters) []Resource {
	// TODO
	// use a segment/selection rule using the assignment system to label and group
	// resources that have been selected
	s.Selected = s.Data

	return s.Selected
}

// Add will add another resource to an entry in data but not selected
func (s *Selection) Add(r Resource) {
	s.Data = append(s.Data, r)
}

// selectors are an internal, simple label type that satisfy the minimum requirement
// for comparison but primarily to aid with yaml serialization/de-serialization. In most
// all cases we will convert to/from labels.Labels
type selectors map[string]string

// ToLabels converts a collection of selectors to the much richer
// labels type from selectors
func (s selectors) ToLabels() labels.Labels {
	l := labels.Labels{}
	for k, v := range s {
		l = append(l, labels.Label{
			Name:  k,
			Value: v,
		})
	}

	return l
}

// FromLabels converts a collection of selectors from the much richer
// labels type to selectors
func FromLabels(l labels.Labels) *selectors {
	s := selectors{}
	for _, label := range l {
		s[label.Name] = label.Value
	}

	return &s
}

func (s selectors) Any() bool {
	return len(s) > 0
}

func (s selectors) Contains(l labels.Label) bool {
	if entry, ok := s[l.Name]; ok {
		return entry == l.Value
	}

	return false
}
