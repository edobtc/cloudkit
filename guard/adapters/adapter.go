package adapters

import (
	"errors"

	"github.com/edobtc/cloudkit/guard/adapters/cloudwatch"
	"github.com/edobtc/cloudkit/guard/adapters/prometheus"

	log "github.com/sirupsen/logrus"
)

var (
	// ErrGuardDoesNotExist is returned if we try to load an adapter
	// that does not exist
	ErrGuardDoesNotExist = errors.New("guard does not exist")
)

// Adapter implements the guard adapter interface
// which allows pluggable guards to be used based on
// experiment plan config
type Adapter interface {
	Check() (bool, error)
}

// Load will load a guard
func Load(name string, config interface{}) (Adapter, error) {
	log.Info(name)

	switch name {
	case "cloudwatch":
		c := config.(cloudwatch.Config)
		return &cloudwatch.Guard{Config: c}, nil
	case "prometheus":
		c := config.(prometheus.Config)
		return &prometheus.Guard{Config: c}, nil
	}

	return nil, ErrGuardDoesNotExist
}
