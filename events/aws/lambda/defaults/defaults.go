package defaults

import (
	"os"

	"github.com/edobtc/cloudkit/config"

	log "github.com/sirupsen/logrus"
)

func SetupDefaults() {
	level := os.Getenv("LOG_LEVEL")

	if level == "" {
		level = "info"
	}

	l, _ := log.ParseLevel(level)
	log.SetLevel(l)

	// Establish configuration defaults
	config.SetLambdaRuntimeDefaults()
}
