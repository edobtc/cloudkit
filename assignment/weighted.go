package assignment

import (
	"errors"
	"strconv"

	"github.com/edobtc/cloudkit/variant"
)

var (
	// ErrFailedAssignmentWeightedChoice returns if there is some
	// inability to select a variant at the set weight
	// when the defaults are returned from Assign
	ErrFailedAssignmentWeightedChoice = errors.New("failed to assign weighted choice")

	// ErrNoVariants Returns when an experiment has no variants
	ErrNoVariants = errors.New("no variants")
)

// WeightedChoice implements an assignment operator
// for weighted choice selection of variants
type WeightedChoice struct {
	Variants []*variant.Variant
}

// NewWeightedChoice returns an initilized weighted choice assignment
// operator
func NewWeightedChoice(v []*variant.Variant) *WeightedChoice {
	return &WeightedChoice{
		Variants: v,
	}
}

// Assign implements assignment to a variant using a weighted choice configuration
func (w *WeightedChoice) Assign(hash uint64) (*variant.Variant, error) {
	scale, _ := strconv.ParseUint("FFFFFFFFFFFFFFF", 16, 64)
	sum := 0.0
	cweights := make([]float64, len(w.Variants))

	if len(w.Variants) == 0 {
		return nil, ErrNoVariants
	}

	for idx, v := range w.Variants {
		sum = sum + v.Weight
		cweights[idx] = sum
	}

	shift := float64(hash) / float64(scale)
	stop := 0.0 + (sum-0.0)*shift

	for idx := range cweights {
		if stop <= cweights[idx] {
			return w.Variants[idx], nil
		}
	}
	return nil, ErrFailedAssignmentWeightedChoice
}
