package sns

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/edobtc/cloudkit/config"
	"github.com/sirupsen/logrus"
)

var (
	ErrNoTopicProvided = errors.New("topic can not be empty")
)

type Publisher struct {
	Topic  string
	Buffer []byte
	svc    *sns.SNS
}

func NewPublisher() *Publisher {
	s := session.Must(session.NewSession())

	return &Publisher{
		Buffer: []byte{},
		svc:    sns.New(s),
		Topic:  config.Read().EventPublisherARN,
	}
}

func (s *Publisher) Send(data []byte) error {
	resp, err := s.svc.Publish(&sns.PublishInput{
		Message:  aws.String(string(data)),
		TopicArn: aws.String(s.Topic),
	})

	logrus.Debug(resp)

	if err != nil {
		return err
	}

	return err
}

func (s *Publisher) Close() error { return nil }
