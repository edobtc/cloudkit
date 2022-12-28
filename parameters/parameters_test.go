package parameters

import (
	"testing"
)

var params = Parameters{
	Parameter{
		Name:  "first",
		Value: 1,
	},
	Parameter{
		Name:  "second",
		Value: 2,
	},
	Parameter{
		Name:  "third",
		Value: 3,
	},
}

func TestWhitelistString(t *testing.T) {
	whitelist := Parameters{
		Parameter{
			Name: "first",
		},
		Parameter{
			Name: "third",
		},
	}

	expected := "first.1.third.3"

	if params.WhitelistedString(whitelist) != expected {
		t.Errorf("expected WhitelistedString() to return %s", expected)
	}
}

func TestStringJoining(t *testing.T) {
	expected := "first.1.second.2.third.3"
	if params.String() != expected {
		t.Errorf("expected String() to return %s", expected)
	}
}

func TestParamsAreOrdered(t *testing.T) {
	unordered := Parameters{
		Parameter{
			Name:  "zlast",
			Value: 1,
		},
		Parameter{
			Name:  "afirst",
			Value: 99,
		},
		Parameter{
			Name:  "middle",
			Value: 9999,
		},
	}

	expected := "afirst.99.middle.9999.zlast.1"

	if unordered.String() != expected {
		t.Errorf("expected String() to return %s", expected)
	}
}
