package keys

import (
	"github.com/edobtc/cloudkit/cli/commands/resources/keys/create"
	"github.com/edobtc/cloudkit/cli/commands/resources/keys/list"
	"github.com/spf13/cobra"
)

// Cmd is the top level relay
// command to export
var Cmd = &cobra.Command{
	Use:   "keys",
	Short: "keys",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	Cmd.AddCommand(create.Cmd)
	Cmd.AddCommand(list.Cmd)
}
