package watchgod

import (
	"github.com/edobtc/cloudkit/cli/commands/watchgod/endpoints"
	"github.com/edobtc/cloudkit/cli/commands/watchgod/publishers"

	"github.com/spf13/cobra"
)

// Cmd is the top level events
// command to export
var Cmd = &cobra.Command{
	Use:   "watchgod",
	Short: "watchgod long running monitor",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	Cmd.AddCommand(publishers.Cmd)
	Cmd.AddCommand(endpoints.Cmd)
}
