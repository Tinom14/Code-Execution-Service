package repository

import (
	"project/processor/api/rabbitMQ/types"
)

type TaskSender interface {
	Send(message types.TaskFromRabbit) error
}
