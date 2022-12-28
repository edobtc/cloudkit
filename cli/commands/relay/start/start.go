package start

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Cmd is the top level relay
// command to export
var Cmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO: relay start")
	},
}
