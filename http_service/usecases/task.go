package usecases

type TaskService interface {
	CreateTask(code string, language string) (string, error)
	GetStatus(taskID string) (string, error)
	SetStatus(taskID, status string) error
	SetResult(taskID string, result string) error
	GetResult(taskID string) (string, error)
}
