package config

import (
	"github.com/edobtc/cloudkit/cli/commands/config/show"

	"github.com/spf13/cobra"
)

// Cmd is the top level config
// command to export
var Cmd = &cobra.Command{
	Use:   "config",
	Short: "config controls everything related to viewing, setting or deleting config values",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	Cmd.AddCommand(show.Cmd)
}
