package start

import (
	"context"

	destinationLoader "github.com/edobtc/cloudkit/events/publishers/autoload"
	sourceLoader "github.com/edobtc/cloudkit/events/subscribers/autoload"
	"github.com/edobtc/cloudkit/logging"
	"github.com/edobtc/cloudkit/multiplex"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	source      string
	destination string
)

func init() {
	logging.Setup()

	Cmd.Flags().StringVarP(&source, "source", "s", "eclair", "source to subscribe to")
	Cmd.Flags().StringVarP(&destination, "destination", "d", "fs", "destination to publish to")
}

var Cmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the multiplexer",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Starting the multiplexer...")
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		s, err := sourceLoader.NewSubscriber(source)
		if err != nil {
			logrus.Error("Failed to create subscriber: ", err)
			return
		}

		d, err := destinationLoader.NewPublisher(destination)
		if err != nil {
			logrus.Error("Failed to create publisher: ", err)
			return
		}

		multiplexer := multiplex.NewMultiplexer(s, d)
		if err := multiplexer.Run(ctx); err != nil {
			logrus.Error("Failed to run multiplexer: ", err)
		}
	},
}
