package events

import (
	// "github.com/edobtc/cloudkit/cli/commands/events/generate"
	"github.com/edobtc/cloudkit/cli/commands/events/inspect"
	"github.com/edobtc/cloudkit/cli/commands/events/subscribe"
	"github.com/edobtc/cloudkit/cli/commands/events/websocket"

	// "github.com/edobtc/cloudkit/cli/commands/events/viewer"

	"github.com/spf13/cobra"
)

// Cmd is the top level events
// command to export
var Cmd = &cobra.Command{
	Use:   "events",
	Short: "events helpers",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	Cmd.AddCommand(subscribe.Cmd)
	Cmd.AddCommand(websocket.Cmd)
	Cmd.AddCommand(inspect.Cmd)
}
