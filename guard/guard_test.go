package guard

import (
	"testing"
)

func TestNothing(t *testing.T) {
	g := NewGuard()

	if g.Interval != DefaultPollInterval {
		t.Errorf("expected %v to equal %v", g.Interval, DefaultPollInterval)
	}
}

func TestDefaults(t *testing.T) {
	g := NewGuard()

	if g.Kind != DefaultKind {
		t.Errorf("expected kind to be Query, got %v", g.Kind)
	}
}

func TestAppendToHistorgram(t *testing.T) {
	g := NewGuard()
	g.WindowSize = 3
	g.StatusDistribution = []Status{
		OK,
		OK,
		OK,
	}

	g.addStatus(Failed)

	if len(g.StatusDistribution) > g.WindowSize {
		t.Error("expected StatusDistribution to not grow larger than WindowSize")
	}

	s := g.StatusDistribution[len(g.StatusDistribution)-1]
	if s != Failed {
		t.Error("expected last element to be of Status Failed")
	}
}
