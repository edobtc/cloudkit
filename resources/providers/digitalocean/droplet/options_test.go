package droplet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRegionName(t *testing.T) {
	type regionTest struct {
		input string
		want  string
		valid bool
	}

	tests := []regionTest{

		{input: "nyc1", want: "New York 1", valid: true},
		{input: "nyc2", want: "New York 2", valid: true},
		{input: "nyc3", want: "New York 3", valid: true},
		{input: "ams1", want: "Amsterdam 1", valid: true},
		{input: "sfo1", want: "San Francisco 1", valid: true},
		{input: "sfo2", want: "San Francisco 2", valid: true},
		{input: "sfo3", want: "San Francisco 3", valid: true},
		{input: "36chambers", want: "Shaolin", valid: false},
	}

	for _, tc := range tests {
		name := GetRegionName(tc.input)
		if tc.valid {
			assert.Equal(t, tc.want, name)
		} else {
			assert.NotEqual(t, tc.want, name)
		}
	}
}

func TestValidRegion(t *testing.T) {
	type regionTest struct {
		input string
		want  bool
	}

	tests := []regionTest{
		{input: "nyc1", want: true},
		{input: "nyc2", want: true},
		{input: "nyc3", want: true},
		{input: "ams1", want: true},
		{input: "sfo1", want: true},
		{input: "sfo2", want: true},
		{input: "sfo3", want: true},
		{input: "36chambers", want: false},
	}

	for _, tc := range tests {
		valid := ValidRegion(tc.input)
		assert.Equal(t, tc.want, valid)
	}
}

func TestGetDropletSize(t *testing.T) {
	type sizeTest struct {
		input string
		want  string
		valid bool
	}

	tests := []sizeTest{
		{input: "default", want: "s-1vcpu-2gb", valid: true},
		{input: "small", want: "s-1vcpu-2gb", valid: true},
		{input: "medium", want: "s-1vcpu-2gb", valid: true},
		{input: "large", want: "s-1vcpu-2gb", valid: true},
	}

	for _, tc := range tests {
		slug := GetDropletSize(tc.input)
		assert.Equal(t, tc.want, slug)
	}

}
