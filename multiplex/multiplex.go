package multiplex

import (
	"context"
	"encoding/json"

	"github.com/edobtc/cloudkit/events/publishers"
	"github.com/edobtc/cloudkit/events/subscribers"
)

type Multiplexer struct {
	source      subscribers.Subscriber
	destination publishers.Publisher

	// config settings
	Behavior Behavior
	Limit    int
}

func NewMultiplexer(
	source subscribers.Subscriber,
	destination publishers.Publisher,
) *Multiplexer {

	cfg := NewDefaultConfig()

	m := &Multiplexer{
		source:      source,
		destination: destination,
	}

	m.Apply(cfg)

	return m
}

func (m *Multiplexer) Apply(cfg *Config) {
	if cfg.Limit != 0 {
		m.Limit = cfg.Limit
	}

	if cfg.Behavior != Unknown {
		m.Behavior = cfg.Behavior
	}
}

func WithLimit(limit int) func(*Multiplexer) *Multiplexer {
	return func(m *Multiplexer) *Multiplexer {
		m.Limit = limit
		return m
	}
}

func WithBehavior(behavior Behavior) func(*Multiplexer) *Multiplexer {
	return func(m *Multiplexer) *Multiplexer {
		m.Behavior = behavior
		return m
	}
}

func WithReadOnly() func(*Multiplexer) *Multiplexer {
	return func(m *Multiplexer) *Multiplexer {
		m.Behavior = Read
		return m
	}
}

func WithWriteOnly() func(*Multiplexer) *Multiplexer {
	return func(m *Multiplexer) *Multiplexer {
		m.Behavior = Write
		return m
	}
}

func WithReadWrite() func(*Multiplexer) *Multiplexer {
	return func(m *Multiplexer) *Multiplexer {
		m.Behavior = ReadWrite
		return m
	}
}

func WithConfig(cfg *Config) func(*Multiplexer) *Multiplexer {
	return func(m *Multiplexer) *Multiplexer {
		m.Apply(cfg)
		return m
	}
}

func WithSource(source subscribers.Subscriber) func(*Multiplexer) *Multiplexer {
	return func(m *Multiplexer) *Multiplexer {
		m.source = source
		return m
	}
}

func WithDestination(destination publishers.Publisher) func(*Multiplexer) *Multiplexer {
	return func(m *Multiplexer) *Multiplexer {
		m.destination = destination
		return m
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

			if m.Behavior != Read {
				if err := m.destination.Send(data); err != nil {
					return err
				}
			}
		case <-wait:
			return nil
		}
	}
}
