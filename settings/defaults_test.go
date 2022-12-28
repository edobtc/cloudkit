package settings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultURL(t *testing.T) {
	assert.Equal(t, "http://127.0.0.1:8081", defaultURL)
}

func TestDefaultSettingsPath(t *testing.T) {
	assert.Equal(t, ".btck/settings.yaml", defaultSettingsPath)
}
