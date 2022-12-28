package webhook

import (
	"bytes"
	"net/http"

	"github.com/edobtc/cloudkit/config"
	"github.com/edobtc/cloudkit/httpclient"

	log "github.com/sirupsen/logrus"
)

type Publisher struct {
	client *http.Client
	URL    string
	Buffer []byte
}

func NewPublisher() (*Publisher, error) {
	return &Publisher{
		Buffer: []byte{},
		client: httpclient.New(),
		URL:    config.Read().EventPublisherName,
	}, nil
}

func (s *Publisher) Send(data []byte) error {
	req, err := http.NewRequest("POST", s.URL, bytes.NewBuffer(s.Buffer))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}

	log.Info("Publisher debug: ", resp.StatusCode)

	return nil
}

func (s *Publisher) Close() error { return nil }
