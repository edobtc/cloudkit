package subscribe

import (
	"fmt"

	"github.com/edobtc/cloudkit/config"
	subscriber "github.com/edobtc/cloudkit/events/subscribers/aws/sqs"

	"github.com/edobtc/cloudkit/cli/commands/events/subscribe/debug"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(debug.Cmd)
}

// Cmd is the command for the cli tool
var Cmd = &cobra.Command{
	Use:   "subscribe",
	Short: "subscribe",
	Run: func(cmd *cobra.Command, args []string) {
		sqsQueueURL := config.Read().EventPublisherSQSQueueURL

		listener, err := subscriber.NewSQSSubscriber(sqsQueueURL)

		if err != nil {
			log.Error(err)
			return
		}

		go func() {
			<-listener.Start()
		}()

		go func() {
			err := <-listener.ErrorChannel
			log.Error(err)
		}()

		for m := range listener.ListenChannel {

			fmt.Println(*m.Body)

			listener.Delete(m.ReceiptHandle)
		}
	},
}
