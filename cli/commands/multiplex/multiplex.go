package multiplex

import (
	"os"

	"github.com/edobtc/cloudkit/cli/commands/multiplex/info"
	"github.com/edobtc/cloudkit/cli/commands/multiplex/listen"
	"github.com/edobtc/cloudkit/cli/commands/multiplex/start"

	"github.com/edobtc/cloudkit/logging"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	source      string
	destination string
)

func init() {
	logging.Setup()

	Cmd.Flags().StringVarP(&source, "source", "s", "", "source to subscribe to")
	Cmd.Flags().StringVarP(&destination, "destination", "d", "", "destination to publish to")
	Cmd.AddCommand(info.Cmd)
	Cmd.AddCommand(listen.Cmd)
	Cmd.AddCommand(start.Cmd)
}

var Cmd = &cobra.Command{
	Use:   "multiplex",
	Short: "multiplex subscribes to events and rebroadcasts on different wires",

	Run: func(cmd *cobra.Command, args []string) {
		log.Info("multiplex")
		cmd.Help()
	},
}

func main() {
	if err := Cmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
