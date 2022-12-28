package relay

import (
	"github.com/edobtc/cloudkit/cli/commands/relay/start"

	"github.com/spf13/cobra"
)

// Cmd is the top level relay
// command to export
var Cmd = &cobra.Command{
	Use:   "relay",
	Short: "relay helpers",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	Cmd.AddCommand(start.Cmd)
}
