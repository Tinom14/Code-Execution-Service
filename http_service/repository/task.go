package repository

type TaskStorage interface {
	CreateTask(taskID string) (string, error)
	GetStatus(taskID string) (string, error)
	SetStatus(taskID string, status string) error
	GetResult(taskID string) (string, error)
	SetResult(taskID string, result string) error
}
