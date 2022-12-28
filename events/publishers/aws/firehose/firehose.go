package firehose

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"

	log "github.com/sirupsen/logrus"
)

type Publisher struct {
	StreamName string
	Buffer     []byte
	svc        *firehose.Firehose
}

func NewPublisher() *Publisher {
	s := session.Must(session.NewSession())

	return &Publisher{
		Buffer: []byte{},
		svc:    firehose.New(s),
	}
}

func (s *Publisher) Send(data []byte) error {
	po, err := s.svc.PutRecord(&firehose.PutRecordInput{
		Record: &firehose.Record{
			Data: data,
		},
		DeliveryStreamName: aws.String(s.StreamName),
	})

	if err != nil {
		return err
	}

	log.Info("Publisher debug: ", po.RecordId)

	return nil
}

func (s *Publisher) Close() error { return nil }
