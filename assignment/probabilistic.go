package assignment

import "github.com/edobtc/cloudkit/variant"

// Probabilistic implements a probabilistic selector of variants
type Probabilistic struct {
	Variants []*variant.Variant
	P        float32
}

// NewProbabilistic returns an initialized Sample operator
func NewProbabilistic(variants []*variant.Variant) *Probabilistic {
	return &Probabilistic{
		Variants: variants,
	}
}

// Assign is not implemented for a Probabilistic Assignment
func (o *Probabilistic) Assign(hash uint64) (*variant.Variant, error) { return nil, nil }
