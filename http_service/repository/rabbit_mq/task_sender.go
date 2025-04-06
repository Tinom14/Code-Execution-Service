package rabbit_mq

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"project/http_service/config"
	"project/processor/api/rabbitMQ/types"
)

type RabbitMQSender struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queueName  string
}

func NewRabbitMQSender(cfg config.RabbitMQ) (*RabbitMQSender, error) {
	url := fmt.Sprintf("amqp://guest:guest@%s:%d", cfg.Host, cfg.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("connecting to rabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		cfg.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQSender{
		connection: conn,
		channel:    ch,
		queueName:  cfg.QueueName,
	}, nil
}

func (r *RabbitMQSender) Send(message types.TaskFromRabbit) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = r.channel.Publish(
		"",
		r.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQSender) Close() {
	r.channel.Close()
	r.connection.Close()
}
