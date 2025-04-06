package postgres

import (
	_ "github.com/lib/pq"
	"project/http_service/api/http/types"
	"project/http_service/repository"
	"project/pkg/postgres_connect"
)

type TaskStorage struct {
	tasks *postgres_connect.PostgresStorage
}

func NewTaskStorage(tasks *postgres_connect.PostgresStorage) *TaskStorage {
	return &TaskStorage{tasks: tasks}
}

func (t *TaskStorage) SendResult(payload types.PostTaskCommitRequest) error {
	_, err := t.tasks.Db.Exec("UPDATE Tasks SET status = 'ready', result = $1 WHERE task_id = $2", payload.Result, payload.TaskID)
	if err != nil {
		return repository.NotFound
	}

	return nil
}
