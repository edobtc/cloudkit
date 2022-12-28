package logging

import (
	"github.com/edobtc/cloudkit/config"

	sd "github.com/joonix/log"
	log "github.com/sirupsen/logrus"
)

func Setup() {
	log.SetFormatter(Formatter())

	lvl, err := log.ParseLevel(config.Read().Verbosity)
	if err != nil {
		log.SetLevel(log.InfoLevel)
		log.Warn(err)
	}
	log.SetLevel(lvl)
}

// Formatter returns a formatter w/ respect to
// a production AppEnvironment being set when
// no explicit formatter is configured
func Formatter() log.Formatter {
	cfg := config.Read()

	// default to text format
	if cfg.LogFormatter == "" {
		return &log.TextFormatter{}
	}

	return loadFormatter(cfg.LogFormatter)
}

// loadFormatter will switch between log formatters
// defaulting to the basic textformatter which we like
// to be text for local environments
func loadFormatter(name string) log.Formatter {
	switch name {
	case "stackdriver", "fluentd", "structured":
		return sd.NewFormatter()
	case "json":
		return &log.JSONFormatter{}
	default:
		return &log.TextFormatter{}
	}
}
