package target

import (
	"testing"
)

func TestHasSelectors(t *testing.T) {
	target := Target{}

	if target.hasSelectors() {
		t.Errorf("expected hasSelectors to be false")
	}

	target = Target{
		Selectors: &selectors{},
	}

	if target.hasSelectors() {
		t.Errorf("expected hasSelectors to be false")
	}

	target = Target{
		Selectors: &selectors{
			"test": "test",
		},
	}

	if target.hasSelectors() == false {
		t.Errorf("expected hasSelectors to be true")
	}
}

func TestSafeNamespace(t *testing.T) {
	target := Target{}

	if target.SafeNamespace() != DefaultNamespace {
		t.Errorf("Expected %s, got %s", DefaultNamespace, target.SafeNamespace())
	}

}
