package endpoints

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Cmd is the command for the cli tool
var Cmd = &cobra.Command{
	Use:   "endpoints",
	Short: "endpoints",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("starting endpoints")
	},
}
