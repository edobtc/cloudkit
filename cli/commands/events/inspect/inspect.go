package inspect

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Cmd is the command for the cli tool
var Cmd = &cobra.Command{
	Use:   "inspect",
	Short: "inspect",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("inspect an event")
		log.Info("add a command to save it")

	},
}
