package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	log "github.com/sirupsen/logrus"
)

const (
	DefaultPollWaitSeconds = 1
	DefaultMaxMessages     = 10
)

// SQSSubscriber is the event subscriber
type SQSSubscriber struct {
	svc                *sqs.SQS
	ListenChannel      chan (*sqs.Message)
	ErrorChannel       chan (error)
	SQSQueueURL        string
	SQSPollWaitSeconds int64
	SQSMaxMessages     int64
	Exit               bool
}

// NewSQSSubscriber returns a new instance of SQSSubscriber
// with default values
func NewSQSSubscriber(url string) (*SQSSubscriber, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	if err != nil {
		return nil, err
	}

	return &SQSSubscriber{
		svc:                sqs.New(sess),
		ListenChannel:      make(chan (*sqs.Message)),
		ErrorChannel:       make(chan (error)),
		SQSQueueURL:        url,
		SQSPollWaitSeconds: DefaultPollWaitSeconds,
		SQSMaxMessages:     DefaultMaxMessages,
	}, nil
}

// Start begins the listener process
func (s *SQSSubscriber) Start() chan (*sqs.Message) {
	go func() {
		for {
			if s.Exit {
				close(s.ListenChannel)
				break
			}

			output, err := s.svc.ReceiveMessage(&sqs.ReceiveMessageInput{
				QueueUrl:            &s.SQSQueueURL,
				MaxNumberOfMessages: aws.Int64(s.SQSMaxMessages),
				WaitTimeSeconds:     aws.Int64(s.SQSPollWaitSeconds),
			})

			if err != nil {
				s.ErrorChannel <- err
				// close(s.ListenChannel)
			}

			for _, message := range output.Messages {
				s.ListenChannel <- message
			}
		}
	}()

	log.Info("starting listener")
	return s.ListenChannel
}

// Delete cleans up and removes messages from the SQS queue
// calling this from the listener handler is required to ensure you
// process each message
func (s *SQSSubscriber) Delete(receiptHandle *string) error {
	_, err := s.svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &s.SQSQueueURL,
		ReceiptHandle: receiptHandle,
	})

	return err
}
