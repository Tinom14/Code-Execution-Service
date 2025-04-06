package usecases

type TaskService interface {
	CompleteTask(taskID string, code string, language string) error
}
