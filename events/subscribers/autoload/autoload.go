package autoload

import (
	"errors"

	"github.com/edobtc/cloudkit/events/subscribers"

	// conforming implementations
	eclair "github.com/edobtc/cloudkit/events/subscribers/lightning/eclair/ws"
	rmq "github.com/edobtc/cloudkit/events/subscribers/rmq"
	zmq "github.com/edobtc/cloudkit/events/subscribers/zmq"
)

// Ensure that the Subscriber interface is properly implemented by the autoloaded subscribers
var _ subscribers.Subscriber = &eclair.EclairSubscriber{}
var _ subscribers.Subscriber = &zmq.Subscriber{}
var _ subscribers.Subscriber = &rmq.RMQSubscriber{}

var (
	ErrAdapterNotFound = errors.New("adapter by name not found")
)

func NewSubscriber(adapter string) (subscribers.Subscriber, error) {
	switch adapter {
	case "eclair":
		return eclair.NewSubscriber(), nil
	case "zmq":
		return zmq.NewSubscriber(), nil
	case "rmq":
		// NewSubscriber returns an err along
		// with the subscriber
		return rmq.NewSubscriber()
	default:
		return nil, ErrAdapterNotFound
	}
}
