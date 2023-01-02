package filesystem

import (
	"fmt"
	"os"

	"github.com/edobtc/cloudkit/config"

	log "github.com/sirupsen/logrus"
)

type Publisher struct {
	Filename string
	Buffer   []byte
	f        *os.File
}

func NewPublisher() (*Publisher, error) {
	f, err := os.Create(fmt.Sprintf("/tmp/%s", config.Read().EventPublisherName))
	if err != nil {
		return nil, err
	}

	return &Publisher{
		Buffer:   []byte{},
		f:        f,
		Filename: config.Read().EventPublisherName,
	}, nil
}

func (s *Publisher) Send(data []byte) error {
	id, err := s.f.Write(append(data, "\n"...))
	if err != nil {
		return err
	}

	s.f.Sync()

	log.Info("Publisher debug: ", id)

	return nil
}

func (s *Publisher) Close() error { return s.f.Close() }
