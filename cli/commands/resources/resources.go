package resources

import (
	"github.com/edobtc/cloudkit/cli/commands/resources/create"
	"github.com/edobtc/cloudkit/cli/commands/resources/list"

	"github.com/spf13/cobra"
)

// Cmd is the top level relay
// command to export
var Cmd = &cobra.Command{
	Use:   "resources",
	Short: "resources",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	Cmd.AddCommand(create.Cmd)
	Cmd.AddCommand(list.Cmd)
}
