package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/edobtc/cloudkit/cmd/multiplex/listen"
	"github.com/edobtc/cloudkit/cmd/multiplex/start"

	"github.com/edobtc/cloudkit/events/publishers/autoload"
	"github.com/edobtc/cloudkit/events/subscribers/zmq"
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

	rootCmd.Flags().StringVarP(&source, "source", "s", "", "source to subscribe to")
	rootCmd.Flags().StringVarP(&destination, "destination", "d", "", "destination to publish to")

	rootCmd.AddCommand(start.Cmd)
	rootCmd.AddCommand(listen.Cmd)
}

var rootCmd = &cobra.Command{
	Use:   "multiplex",
	Short: "multiplex subscribes to events and rebroadcasts on different wires",

	Run: func(cmd *cobra.Command, args []string) {
		log.Info("starting")

		srv := zmq.NewSubscriber()
		log.Info(srv)

		signal.Notify(srv.Close, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

		// // move these to config
		srv.Subscribe("hashtx")
		srv.Subscribe("hashblock")

		fs, err := autoload.NewPublisher("fs")
		if err != nil {
			log.Error(err)
			return
		}

		srv.AddDestination(fs)

		go func() {
			s := <-srv.Close
			log.Infof("received signal %v", s)
			srv.Detach()
		}()

		<-srv.Start()

	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
