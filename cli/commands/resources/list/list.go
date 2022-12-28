package list

import (
	"github.com/edobtc/cloudkit/cli/commands/resources/list/accounts"

	"github.com/spf13/cobra"
)

// Cmd is the top level relay
// command to export
var Cmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	Cmd.AddCommand(accounts.Cmd)
}
