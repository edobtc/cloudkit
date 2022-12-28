package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	// Services

	sink "github.com/edobtc/cloudkit/events/sink/http"
	relay "github.com/edobtc/cloudkit/relay/http"

	// middelwares

	"github.com/edobtc/cloudkit/http/middleware/namespace"

	"github.com/edobtc/cloudkit/logging"
	"github.com/edobtc/cloudkit/server"
)

func init() {
	logging.Setup()
}

var rootCmd = &cobra.Command{
	Use:   "btck",
	Short: "btck bundles the entire package",
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.NewServer()

		r := srv.Router()

		// middleware attachment
		r.Use(namespace.DecorateNamespace)

		relay.BindRoutesWithPrefix(r)
		sink.BindRoutesWithPrefix(r)

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
