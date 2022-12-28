package main

import (
	"github.com/edobtc/cloudkit/events/aws/lambda/defaults"
	"github.com/edobtc/cloudkit/events/handlers/bitcoind/rpc"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	TxID string
}

func main() {
	defaults.SetupDefaults()
	lambda.Start(rpc.GetTransaction)
}
