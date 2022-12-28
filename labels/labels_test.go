package labels

import (
	"os"
	"testing"
)

func TestLabelNameDefaults(t *testing.T) {
	if name := ExperimentLabelName(); name != ExperimentLabelDefault {
		t.Errorf("expected %s to eq %s", ExperimentLabelDefault, name)
	}

	if name := AssignmentLabelName(); name != AssignmentLabelDefault {
		t.Errorf("expected %s to eq %s", AssignmentLabelDefault, name)
	}

	if name := VariantLabelName(); name != VariantLabelDefault {
		t.Errorf("expected %s to eq %s", VariantLabelDefault, name)
	}
}

func TestLabelNameOverrides(t *testing.T) {
	expOverrideValue := "EXPERIMENT_OVERRIDE"
	os.Setenv(experimentLabelKey, expOverrideValue)

	if name := ExperimentLabelName(); name != expOverrideValue {
		t.Errorf("expected %s to eq %s", expOverrideValue, name)
	}

	asOverrideValue := "ASSIGNMENT_OVERRIDE"
	os.Setenv(assignmentLabelKey, asOverrideValue)
	if name := AssignmentLabelName(); name != asOverrideValue {
		t.Errorf("expected %s to eq %s", asOverrideValue, name)
	}

	vrOverrideValue := "VARIANT_OVERRIDE"
	os.Setenv(variantLabelKey, vrOverrideValue)
	if name := VariantLabelName(); name != vrOverrideValue {
		t.Errorf("expected %s to eq %s", vrOverrideValue, name)
	}
}
