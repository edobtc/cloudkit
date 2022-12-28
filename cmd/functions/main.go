package main

import (
	"context"
	"flag"

	"github.com/edobtc/cloudkit/events/handlers/bitcoind/rpc"

	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

var (
	handlerName string
)

func init() {
	flag.StringVar(&handlerName, "handler", "status", "the handler name to execute")

	flag.Parse()

	// Establish Lambda Defaults
	// defaults.SetupDefaults()
}

func main() {
	log.Info(handlerName)

	switch handlerName {
	case "getblock":
		log.Info("running rpc getblock")
		lambda.Start(rpc.GetBlock)
	case "basic":
		log.Info("running rpc GetTransaction")
		lambda.Start(rpc.GetTransaction)
	default:
		log.Info("handler name not found defaulting to status")
		lambda.Start(StatusHandler)
	}
}

// Status Handler Setup
type Event struct{}

func StatusHandler(ctx context.Context, event Event) (string, error) {
	return "ok", nil
}
