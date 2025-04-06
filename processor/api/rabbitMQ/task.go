package rabbitMQ

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"project/processor/api/rabbitMQ/types"
	"project/processor/config"
	"project/processor/usecases"
)

type RabbitMQReceiver struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queueName  string
	service    usecases.TaskService
}

func NewRabbitMQReceiver(cfg config.RabbitMQ, service usecases.TaskService) (*RabbitMQReceiver, error) {
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
	return &RabbitMQReceiver{
		connection: conn,
		channel:    ch,
		queueName:  cfg.QueueName,
		service:    service,
	}, nil
}

func (r *RabbitMQReceiver) Receive() error {
	log.Printf("CodeProcessor started. Waiting for messages...")
	messages, err := r.channel.Consume(
		r.queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Errorf("consuming RabbitMQ messages: %w", err)
	}
	for message := range messages {
		var task types.TaskFromRabbit
		if err := json.Unmarshal(message.Body, &task); err != nil {
			log.Printf("failed to unmarshal task: %v", err)
			continue
		}
		if err := r.service.CompleteTask(task.TaskID, task.Language, task.Code); err != nil {
			log.Printf("failed to complete task: %v", err)
			continue
		}
	}
	return nil
}

func (r *RabbitMQReceiver) Close() {
	r.channel.Close()
	r.connection.Close()
}
