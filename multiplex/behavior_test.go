package multiplex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBehaviorString(t *testing.T) {
	tests := []struct {
		behavior Behavior
		expected string
	}{
		{ReadWrite, "ReadWrite"},
		{Read, "Read"},
		{Write, "Write"},
		{Behavior(0), "Unknown"},
	}

	for _, test := range tests {
		result := test.behavior.String()
		assert.Equal(t, test.expected, result, "Unexpected result for behavior %v", test.behavior)
	}
}

func TestBehaviorFromString(t *testing.T) {
	tests := []struct {
		str      string
		expected Behavior
	}{
		{"ReadWrite", ReadWrite},
		{"Read", Read},
		{"Write", Write},
		{"Unknown", Unknown},
	}

	for _, test := range tests {
		result, err := BehaviorFromString(test.str)
		assert.NoError(t, err, "Unexpected error for string %s", test.str)
		assert.Equal(t, test.expected, result, "Unexpected result for string %s", test.str)
	}
}
