package batch

import (
	"fmt"

	"github.com/edobtc/cloudkit/events/formats"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	eventCount = 3000
)

// Cmd is the command for the cli tool
var Cmd = &cobra.Command{
	Use:   "batch",
	Short: "batch",
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < eventCount; i++ {
			log.Info("generate an event")

			e := formats.GenericEvent{
				Name: "test event",
			}

			fmt.Println(e)
		}
	},
}
