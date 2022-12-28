package generate

import (
	"github.com/edobtc/cloudkit/cli/commands/events/generate/batch"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Cmd is the command for the cli tool
var Cmd = &cobra.Command{
	Use:   "generate",
	Short: "generate",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("generating a single event")
		log.Info("add a command to save it")

	},
}

func init() {
	Cmd.AddCommand(batch.Cmd)
}
