package variant

import (
	"testing"
)

func TestAutoLabelDefault(t *testing.T) {
	v := NewVariant()
	if v.AutoLabel == false {
		t.Errorf("expected AutoLabel to be true")
	}
}
