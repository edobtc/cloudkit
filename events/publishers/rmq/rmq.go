package rmq

import (
	"fmt"

	"github.com/edobtc/cloudkit/config"
	"github.com/streadway/amqp"
)

type Publisher struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queueName  string
}

func NewPublisher() (*Publisher, error) {
	queueName := config.Read().RabbitMQ.QueueName

	fmt.Println(config.Read().RabbitMQ.URL)

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
		config.Read().RabbitMQ.Durable,    // Durable
		config.Read().RabbitMQ.AutoDelete, // Delete when unused
		config.Read().RabbitMQ.Exclusive,  // Exclusive
		config.Read().RabbitMQ.NoWait,     // No-wait
		nil,                               // Arguments
	)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		connection: conn,
		channel:    ch,
		queueName:  queueName,
	}, nil
}

func (r *Publisher) Listen() (<-chan amqp.Delivery, error) {
	queueName := config.Read().RabbitMQ.QueueName

	msgs, err := r.channel.Consume(
		queueName,                        // queue
		"",                               // consumer
		true,                             //config.Read().RabbitMQ.AutoAck,  // auto-ack
		config.Read().RabbitMQ.Exclusive, // exclusive
		false,                            // no-local
		config.Read().RabbitMQ.NoWait,    // no-wait
		nil,                              // args
	)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (r *Publisher) Send(data []byte) error {
	return r.channel.Publish(
		config.Read().RabbitMQ.ExchangeName, // exchange
		config.Read().RabbitMQ.QueueName,    // routing key
		config.Read().RabbitMQ.Mandatory,    // mandatory
		config.Read().RabbitMQ.Immediate,    // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  config.Read().RabbitMQ.ContentType,
			Body:         data,
		},
	)
}

func (r *Publisher) Close() error {
	if err := r.channel.Close(); err != nil {
		return err
	}
	return r.connection.Close()
}
