package ws

import (
	"os"

	eclair "github.com/edobtc/go-eclair"
)

type EclairSubscriber struct {
	client   *eclair.Client
	internal chan (*eclair.Message)

	listener chan interface{}
	Close    chan os.Signal
	Wait     chan bool
	Error    chan error
}

func NewSubscriber() *EclairSubscriber {
	c := eclair.NewClient()

	return &EclairSubscriber{
		client: c,

		listener: make(chan interface{}),
		Close:    make(chan os.Signal),
		Wait:     make(chan bool),
		Error:    make(chan error),
	}
}

func (s *EclairSubscriber) Start() chan bool {
	channel, err := s.client.Subscribe()
	if err != nil {
		return s.Wait
	}

	go func() {
		for {
			select {
			case msg := <-channel:
				s.listener <- msg
			case <-s.Close:
				s.Wait <- true
				return
			}
		}
	}()

	return s.Wait
}

func (s *EclairSubscriber) Detach() error {
	close(s.Close)
	return nil
}

func (s *EclairSubscriber) Listen() <-chan interface{} {
	return s.listener
}
