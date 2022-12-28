package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigPreparedOnlyOnce(t *testing.T) {
	updatedValue := "wutang"

	cfg := Read()

	os.Setenv("AUTH_SECRET", updatedValue)

	assert.NotEqual(t, cfg.Verbosity, updatedValue, "they should not be equal")
}
