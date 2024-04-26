package main

import (
	"os"

	mx "github.com/edobtc/cloudkit/cli/commands/multiplex"

	log "github.com/sirupsen/logrus"
)

var rootCmd = mx.Cmd

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
