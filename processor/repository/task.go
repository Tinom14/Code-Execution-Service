package repository

import (
	"project/http_service/api/http/types"
)

type TaskSender interface {
	SendResult(task types.PostTaskCommitRequest) error
}
