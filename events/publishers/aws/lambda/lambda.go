package lambda

import (
	"encoding/json"

	"github.com/edobtc/cloudkit/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"

	log "github.com/sirupsen/logrus"
)

var ApplicationPublisher = NewPublisher()

type Publisher struct {
	FunctionArn string
	Buffer      []byte
	svc         *lambda.Lambda
}

func NewPublisher() *Publisher {
	s := session.Must(session.NewSession())

	return &Publisher{
		Buffer:      []byte{},
		svc:         lambda.New(s),
		FunctionArn: config.Read().EventPublisherName,
	}
}

func (s *Publisher) Send(data []byte) error {
	log.Info("transmitting data")
	log.Info(string(data))

	payload := Event{
		Name: "publisherHandler",
		Data: string(data),
	}

	p, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	po, err := s.svc.Invoke(&lambda.InvokeInput{
		Payload:      p,
		FunctionName: aws.String(s.FunctionArn),
	})

	if err != nil {
		return err
	}

	log.Info("Publisher debug: ", po.GoString())

	return nil
}

func (s *Publisher) Close() error { return nil }
