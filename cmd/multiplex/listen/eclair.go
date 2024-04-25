package listen

import (
	"github.com/edobtc/cloudkit/events/subscribers/autoload"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var CmdEclair = &cobra.Command{
	Use:   "listen",
	Short: "Listens on the specified source",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Listening on source: ", source)
		subscriber, err := autoload.NewSubscriber("eclair")
		if err != nil {
			logrus.Error("Failed to create subscriber: ", err)
			return
		}
		subscriber.Start()

		for msg := range subscriber.Listen() {
			logrus.Info("Received message: ", msg)
		}
	},
}
