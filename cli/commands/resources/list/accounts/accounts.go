package accounts

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "accounts",
	Short: "provision accounts storage",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("list stuff")
	},
}
