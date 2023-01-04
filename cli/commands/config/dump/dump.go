package dump

import (
	"fmt"

	"github.com/edobtc/cloudkit/config"
	"github.com/spf13/cobra"
)

// Cmd is the command for the cli tool
var Cmd = &cobra.Command{
	Use:   "dump",
	Short: "dump",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Read()

		fmt.Println(cfg)
	},
}
