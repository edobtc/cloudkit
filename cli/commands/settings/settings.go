package settings

import (
	"github.com/edobtc/cloudkit/cli/commands/settings/reset"
	"github.com/edobtc/cloudkit/cli/commands/settings/set"
	"github.com/edobtc/cloudkit/cli/commands/settings/show"

	"github.com/spf13/cobra"
)

// Cmd is the top level settings
// command to export
var Cmd = &cobra.Command{
	Use:   "settings",
	Short: "settings controls everything related to viewing, setting or deleting settings values",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	Cmd.AddCommand(reset.Cmd)
	Cmd.AddCommand(set.Cmd)
	Cmd.AddCommand(show.Cmd)

	// TODO
	// add a set
	// add a delete
}
