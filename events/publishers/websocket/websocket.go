package websocket

import (
	"github.com/edobtc/cloudkit/controlplane/ws"
)

type Publisher struct {
}

func NewPublisher() (*Publisher, error) {
	return &Publisher{}, nil
}

func (s *Publisher) Send(data []byte) error {
	// this is greedy and just swallows everything
	// at the moment, is 'best effort' i suppose
	//
	// https://www.youtube.com/watch?v=8FIBPKRV3kk
	ws.Pool.Publish(data)
	// this needs to be better
	return nil
}

func (s *Publisher) Close() error { return nil }
