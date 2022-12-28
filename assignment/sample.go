package assignment

import "github.com/edobtc/cloudkit/variant"

// Sample implements a sampler for
type Sample struct {
	Variants []*variant.Variant
}

// NewSample returns an initialized Sample operator
func NewSample(variants []*variant.Variant) *Sample {
	return &Sample{
		Variants: variants,
	}
}

// Assign is not implemented
func (o *Sample) Assign(hash uint64) (*variant.Variant, error) { return nil, nil }
