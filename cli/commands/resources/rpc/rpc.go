package rpc

import (
	"github.com/edobtc/cloudkit/cli/commands/resources/rpc/credentials"
	"github.com/spf13/cobra"
)

// Cmd is the top level relay
// command to export
var Cmd = &cobra.Command{
	Use:   "rpc",
	Short: "rpc",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	Cmd.AddCommand(credentials.Cmd)
}
