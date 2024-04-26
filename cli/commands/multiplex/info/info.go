package info

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "info",
	Short: "Prints the multiplexer info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("multiplexer info...")
	},
}
