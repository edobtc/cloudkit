package main

import (
	"os"

	"github.com/edobtc/cloudkit/cmd/multiplex/info"
	"github.com/edobtc/cloudkit/cmd/multiplex/listen"
	"github.com/edobtc/cloudkit/cmd/multiplex/start"

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

	rootCmd.AddCommand(info.Cmd)
	rootCmd.AddCommand(listen.Cmd)
	rootCmd.AddCommand(start.Cmd)
}

var rootCmd = &cobra.Command{
	Use:   "multiplex",
	Short: "multiplex subscribes to events and rebroadcasts on different wires",

	Run: func(cmd *cobra.Command, args []string) {
		log.Info("multiplex")
		cmd.Help()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
