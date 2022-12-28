package assignment

import "github.com/edobtc/cloudkit/variant"

// UniformChoice implements an assignment operator
// for uniform choice selection of variants
type UniformChoice struct {
	Variants []*variant.Variant
}

// NewUniformChoice returns an initilized weighted choice assignment
// operator
func NewUniformChoice(v []*variant.Variant) *UniformChoice {
	return &UniformChoice{
		Variants: v,
	}
}

// Assign implements assignment to a variant using a uniform choice configuration
func (u *UniformChoice) Assign(hash uint64) (*variant.Variant, error) {
	idx := hash % uint64(len(u.Variants))
	return u.Variants[idx], nil
}
