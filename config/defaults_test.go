package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultProvider(t *testing.T) {
	assert.Equal(t, DefaultNodeHost, "127.0.0.1:8332", "they should be equal")
}

func TestDefaultSSHKeyName(t *testing.T) {
	assert.Equal(t, DefaultSSHKeyName, "cloudkit-ssh-key", "they should be equal")
}
