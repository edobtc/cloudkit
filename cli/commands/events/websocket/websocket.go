package websocket

import (
	"github.com/edobtc/cloudkit/cli/commands/events/websocket/publish"
	"github.com/edobtc/cloudkit/cli/commands/events/websocket/subscribe"
	"github.com/spf13/cobra"
)

var (
	host string
)

func init() {
	Cmd.PersistentFlags().StringVarP(&host, "host", "n", "127.0.0.1:8081", "host")
	Cmd.AddCommand(subscribe.Cmd)
	Cmd.AddCommand(publish.Cmd)
}

// Cmd is the command for the cli tool
var Cmd = &cobra.Command{
	Use:   "websocket",
	Short: "websocket",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
