package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/edobtc/cloudkit/logging"
	"github.com/edobtc/cloudkit/server"

	// relay route setup
	gateway "github.com/edobtc/cloudkit/gateway/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	logging.Setup()
}

var rootCmd = &cobra.Command{
	Use:   "cloudkit-gateway",
	Short: "cloudkit-gateway runs a server for intercepting requests to nodes and providing some facility to audit events, add rules etc",

	Run: func(cmd *cobra.Command, args []string) {
		srv := server.NewServer()

		r := srv.Router()

		// Account View Handlers
		gateway.BindRoutes(r)

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
