package rmq

import (
	"github.com/edobtc/cloudkit/config"
	"github.com/streadway/amqp"
)

type RMQSubscriber struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queueName  string
}

func NewSubscriber() (*RMQSubscriber, error) {
	queueName := config.Read().RabbitMQ.QueueName

	conn, err := amqp.Dial(config.Read().RabbitMQ.URL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		queueName,
		config.Read().RabbitMQ.Durable,
		config.Read().RabbitMQ.AutoDelete,
		config.Read().RabbitMQ.Exclusive,
		config.Read().RabbitMQ.NoWait,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &RMQSubscriber{
		connection: conn,
		channel:    ch,
		queueName:  queueName,
	}, nil
}

func (r *RMQSubscriber) Start() chan bool {
	// Start consuming messages
	msgs, err := r.channel.Consume(
		r.queueName,
		"",
		config.Read().RabbitMQ.AutoAck,
		config.Read().RabbitMQ.Exclusive,
		false,
		config.Read().RabbitMQ.NoWait,
		nil,
	)
	if err != nil {
		panic(err)
	}

	done := make(chan bool)
	go func() {
		for d := range msgs {
			// Process message
			// Placeholder: Print message to console
			println("Received message: ", string(d.Body))
		}
		done <- true
	}()
	return done
}

func (r *RMQSubscriber) Detach() error {
	if err := r.channel.Close(); err != nil {
		return err
	}
	return r.connection.Close()
}

func (r *RMQSubscriber) Listen() <-chan interface{} {
	msgs, err := r.channel.Consume(
		r.queueName,
		"",
		config.Read().RabbitMQ.AutoAck,
		config.Read().RabbitMQ.Exclusive,
		false,
		config.Read().RabbitMQ.NoWait,
		nil,
	)
	if err != nil {
		panic(err)
	}

	output := make(chan interface{})
	go func() {
		for d := range msgs {
			output <- d.Body
		}
		close(output)
	}()
	return output
}
