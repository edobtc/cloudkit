package listen

import (
	"github.com/edobtc/cloudkit/events/subscribers/autoload"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var CmdZmq = &cobra.Command{
	Use:   "zmq",
	Short: "Listens on the specified zmq source",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Listening on source: ", source)
		subscriber, err := autoload.NewSubscriber("zmq")
		if err != nil {
			logrus.Error("Failed to create subscriber: ", err)
			return
		}

		subscriber.Start()

		for msg := range subscriber.Listen() {
			logrus.Info("Received message: ", msg)
			logrus.Info("Publishing message to destination: ", destination)
		}
	},
}
