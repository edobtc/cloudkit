package kinesis

import (
	"time"

	"github.com/edobtc/cloudkit/aws/session"
	"github.com/edobtc/cloudkit/config"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/kinesis"

	log "github.com/sirupsen/logrus"
)

type Publisher struct {
	Name   string
	Buffer []byte
	svc    *kinesis.Kinesis
}

func NewPublisher() *Publisher {
	s := session.NewDynamicSession()

	return &Publisher{
		Buffer: []byte{},
		svc:    kinesis.New(s),
		Name:   config.Read().EventPublisherName,
	}
}

func (s *Publisher) Send(data []byte) error {
	po, err := s.svc.PutRecord(&kinesis.PutRecordInput{
		Data:         data,
		StreamName:   aws.String(s.Name),
		PartitionKey: aws.String(time.Now().String()),
	})

	if err != nil {
		return err
	}
	log.Debug("Publisher debug: ", po.SequenceNumber)

	return nil
}

func (s *Publisher) Transmit() error {
	po, err := s.svc.PutRecord(&kinesis.PutRecordInput{
		StreamName:   aws.String(s.Name),
		PartitionKey: aws.String("key1"),
	})

	if err != nil {
		return err
	}

	log.Debug("Publisher debug: ", po.SequenceNumber)

	return nil
}

func (s *Publisher) Close() error { return nil }
