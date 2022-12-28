package autoload

import (
	"testing"
)

func TestExists(t *testing.T) {
	if Exists("whatever/bloop") {
		t.Error("Exists should not find whatever/bloop")
	}

	if Exists("aws/lambda") == false {
		t.Error("Exists should find aws/lambda")
	}
}
