package ram_storage

import (
	"project/http_service/domain"
	"project/http_service/repository"
	"sync"
)

type TaskRamStorage struct {
	mu    sync.Mutex
	tasks map[string]domain.Task
}

func NewTaskRamStorage() *TaskRamStorage {
	return &TaskRamStorage{
		tasks: make(map[string]domain.Task),
	}
}

func (s *TaskRamStorage) CreateTask(taskID string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks[taskID] = domain.Task{Status: "in_progress"}
	return taskID, nil
}

func (s *TaskRamStorage) GetStatus(taskID string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[taskID]
	if !exists {
		return "", repository.NotFound
	}
	return task.Status, nil
}

func (s *TaskRamStorage) SetStatus(taskID string, status string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[taskID]
	if !exists {
		return repository.NotFound
	}
	task.Status = status
	s.tasks[taskID] = task
	return nil
}

func (s *TaskRamStorage) GetResult(taskID string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[taskID]
	if !exists {
		return "", repository.NotFound
	}
	return task.Result, nil
}

func (s *TaskRamStorage) SetResult(taskID string, result string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[taskID]
	if !exists {
		return repository.NotFound
	}
	task.Result = result
	s.tasks[taskID] = task
	return nil
}
