package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/edobtc/cloudkit/events/publishers/autoload"
	"github.com/edobtc/cloudkit/events/subscribers/zmq"
	"github.com/edobtc/cloudkit/logging"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	logging.Setup()
}

var rootCmd = &cobra.Command{
	Use:   "multiplex",
	Short: "multiplex subscribes to events and rebroadcasts on different wires",

	Run: func(cmd *cobra.Command, args []string) {
		log.Info("starting")
		srv := zmq.New()
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
