package multiplex

import (
	"context"
	"encoding/json"

	"github.com/edobtc/cloudkit/events/publishers"
	"github.com/edobtc/cloudkit/events/subscribers"
)

type Config struct {
	// Limit is the maximum number of messages to send
	// to the destination publisher before ending
	Limit int `json:"limit"`
}

type Multiplexer struct {
	source      subscribers.Subscriber
	destination publishers.Publisher
}

func NewMultiplexer(source subscribers.Subscriber, destination publishers.Publisher) *Multiplexer {
	return &Multiplexer{
		source:      source,
		destination: destination,
	}
}

func WithSource(source subscribers.Subscriber) func(*Multiplexer) {
	return func(m *Multiplexer) {
		m.source = source
	}
}

func WithDestination(destination publishers.Publisher) func(*Multiplexer) {
	return func(m *Multiplexer) {
		m.destination = destination
	}
}

func (m *Multiplexer) Run(ctx context.Context) error {
	wait := m.source.Start()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-m.source.Listen():
			data, err := json.Marshal(msg)
			if err != nil {
				return ErrMarshallingMessage
			}

			if err := m.destination.Send(data); err != nil {
				return err
			}
		case <-wait:
			return nil
		}
	}
}
