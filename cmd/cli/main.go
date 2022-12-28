package main

import (
	"os"

	"github.com/edobtc/cloudkit/cli/commands/about"
	"github.com/edobtc/cloudkit/cli/commands/config"
	"github.com/edobtc/cloudkit/cli/commands/events"
	"github.com/edobtc/cloudkit/cli/commands/relay"
	"github.com/edobtc/cloudkit/cli/commands/resources"
	"github.com/edobtc/cloudkit/cli/commands/resources/keys"
	"github.com/edobtc/cloudkit/cli/commands/settings"
	"github.com/edobtc/cloudkit/cli/commands/status"
	"github.com/edobtc/cloudkit/version"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(formatter())

	level := os.Getenv("LOG_LEVEL")

	if level == "" {
		level = "info"
	}

	l, _ := log.ParseLevel(level)
	log.SetLevel(l)

	bootstrap()
}

var bck = &cobra.Command{
	Use:     "bck",
	Version: version.String(),
	Short:   "bck",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func bootstrap() {
	// Add commands here
	bck.AddCommand(about.Cmd)
	bck.AddCommand(config.Cmd)
	bck.AddCommand(events.Cmd)
	bck.AddCommand(keys.Cmd)
	bck.AddCommand(relay.Cmd)
	bck.AddCommand(settings.Cmd)
	bck.AddCommand(status.Cmd)

	// resource commands
	bck.AddCommand(resources.Cmd)
}

func main() {
	if err := bck.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func formatter() log.Formatter {
	if format := os.Getenv("LOG_FORMAT"); format == "JSON" {
		return &log.JSONFormatter{}
	}

	return &log.TextFormatter{}
}
