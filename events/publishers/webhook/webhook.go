package webhook

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/edobtc/cloudkit/config"
	"github.com/edobtc/cloudkit/httpclient"
	"github.com/edobtc/cloudkit/version"

	log "github.com/sirupsen/logrus"
)

var (
	userAgent = fmt.Sprintf("edobtc-cloudkit-webhook-notifications/%s", version.Version)
)

type Publisher struct {
	client *http.Client
	URL    string
}

func NewPublisher() (*Publisher, error) {
	return &Publisher{
		client: httpclient.New(),
		URL:    config.Read().Notifications.WebhookURL,
	}, nil
}

func (s *Publisher) Send(data []byte) error {
	req, err := http.NewRequest("POST", s.URL, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}

	log.Info("webhook http status: ", resp.StatusCode)

	return nil
}

func (s *Publisher) Close() error { return nil }
