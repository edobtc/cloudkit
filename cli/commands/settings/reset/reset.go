package reset

import (
	"github.com/edobtc/cloudkit/settings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Cmd is the command for the cli tool
var Cmd = &cobra.Command{
	Use:   "reset",
	Short: "reset",
	Run: func(cmd *cobra.Command, args []string) {
		if err := settings.Reset(); err != nil {
			log.Error("failed to reset settings")
			log.Error(err)
			return
		}

		log.Info("settings reset")
	},
}
