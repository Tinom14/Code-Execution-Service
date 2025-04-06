package postgres

import (
	"fmt"
	_ "github.com/lib/pq"
	"project/http_service/repository"
	"project/pkg/postgres_connect"
)

type TaskStorage struct {
	tasks *postgres_connect.PostgresStorage
}

func NewTaskStorage(tasks *postgres_connect.PostgresStorage) *TaskStorage {
	return &TaskStorage{tasks: tasks}
}

func (t *TaskStorage) CreateTask(taskID string) (string, error) {
	_, err := t.tasks.Db.Exec("INSERT INTO Tasks (task_id, status) VALUES ($1, 'in_progress')", taskID)
	if err != nil {
		return "", fmt.Errorf("failed to create task: %v", err)
	}
	return taskID, nil
}

func (t *TaskStorage) GetStatus(taskID string) (string, error) {
	row := t.tasks.Db.QueryRow("SELECT status FROM Tasks WHERE task_id = $1", taskID)
	var status string
	err := row.Scan(&status)
	if err != nil {
		return "", repository.NotFound
	}
	return status, nil
}

func (t *TaskStorage) SetStatus(taskID string, status string) error {
	_, err := t.tasks.Db.Exec("UPDATE Tasks SET status = $1 WHERE task_id = $2", status, taskID)
	if err != nil {
		return repository.NotFound
	}
	return nil
}

func (t *TaskStorage) GetResult(taskID string) (string, error) {
	row := t.tasks.Db.QueryRow("SELECT result FROM Tasks WHERE task_id = $1", taskID)
	var result string
	err := row.Scan(&result)
	if err != nil {
		return "", repository.NotFound
	}
	return result, nil
}

func (t *TaskStorage) SetResult(taskID string, result string) error {
	_, err := t.tasks.Db.Exec("UPDATE Tasks SET result = $1 WHERE task_id = $2", result, taskID)
	if err != nil {
		return repository.NotFound
	}
	return nil
}
