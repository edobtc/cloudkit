package autoload

import (
	"errors"

	delivery "github.com/edobtc/cloudkit/events/publishers"

	// conforming implementations
	"github.com/edobtc/cloudkit/events/publishers/aws/firehose"
	"github.com/edobtc/cloudkit/events/publishers/aws/kinesis"
	"github.com/edobtc/cloudkit/events/publishers/aws/lambda"
	"github.com/edobtc/cloudkit/events/publishers/aws/sns"
	"github.com/edobtc/cloudkit/events/publishers/filesystem"
	"github.com/edobtc/cloudkit/events/publishers/webhook"
	"github.com/edobtc/cloudkit/events/publishers/websocket"
)

var (
	ErrAdapterNotFound = errors.New("adapter by name not found")
)

// NewPublisher loads a Publisher implementation
func NewPublisher(adapter string) (delivery.Publisher, error) {
	switch adapter {
	case "firehose":
		return firehose.NewPublisher(), nil
	case "sns":
		return sns.NewPublisher(), nil
	case "lambda":
		return lambda.NewPublisher(), nil
	case "kinesis":
		return kinesis.NewPublisher(), nil
	case "filesystem", "fs":
		return filesystem.NewPublisher()
	case "webhook", "wh":
		return webhook.NewPublisher()
	case "websocket", "ws":
		return websocket.NewPublisher()
	default:
		return nil, ErrAdapterNotFound
	}
}
