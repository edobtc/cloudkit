package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/edobtc/cloudkit/events/aws/lambda/defaults"
	log "github.com/sirupsen/logrus"
)

func main() {
	defaults.SetupDefaults()

	log.Info("running status")
	lambda.Start(StatusHandler)
}

// Status Handler Setup
type Event struct{}

func StatusHandler(ctx context.Context, event Event) (string, error) {
	return "ok", nil
}
