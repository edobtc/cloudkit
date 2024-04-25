package listen

import (
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

	CmdEclair.Flags().StringVarP(&source, "source", "s", "", "source to subscribe to")
	CmdEclair.Flags().StringVarP(&destination, "destination", "d", "", "destination to publish to")

	Cmd.AddCommand(CmdZmq)
	Cmd.AddCommand(CmdEclair)
}

var Cmd = &cobra.Command{
	Use:   "listen",
	Short: "listen",

	Run: func(cmd *cobra.Command, args []string) {
		log.Info("listen")

	},
}
