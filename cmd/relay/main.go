package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/edobtc/cloudkit/logging"
	"github.com/edobtc/cloudkit/server"

	// relay route setup
	relay "github.com/edobtc/cloudkit/relay/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	logging.Setup()
}

var rootCmd = &cobra.Command{
	Use:   "cloudkit-relay",
	Short: "cloudkit-relay runs a handler for collecting events event streams and batch processing them",

	Run: func(cmd *cobra.Command, args []string) {
		srv := server.NewServer()

		r := srv.Router()

		// Account View Handlers
		relay.BindRoutes(r)

		signal.Notify(srv.Close, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

		go func() {
			s := <-srv.Close
			log.Infof("received signal %v", s)
			srv.GracefulShutdown()
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
