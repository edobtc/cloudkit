package websocket

import (
	"github.com/edobtc/cloudkit/cli/commands/events/websocket/subscribe"
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(subscribe.Cmd)
}

// Cmd is the command for the cli tool
var Cmd = &cobra.Command{
	Use:   "websocket",
	Short: "websocket",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
