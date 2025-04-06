package service

import (
	"fmt"
	"github.com/google/uuid"
	"project/http_service/repository"
	"project/processor/api/rabbitMQ/types"
)

type Task struct {
	repo   repository.TaskStorage
	sender repository.TaskSender
}

func NewTaskService(repo repository.TaskStorage, sender repository.TaskSender) *Task {
	return &Task{
		repo:   repo,
		sender: sender,
	}
}

func (s *Task) CreateTask(code string, language string) (string, error) {
	var taskID = uuid.New().String()
	message := types.TaskFromRabbit{
		TaskID:   taskID,
		Code:     code,
		Language: language,
	}
	err := s.sender.Send(message)
	if err != nil {
		return "", fmt.Errorf("sending object: %w", err)
	}
	return s.repo.CreateTask(taskID)
}

func (s *Task) GetStatus(taskID string) (string, error) {
	return s.repo.GetStatus(taskID)
}

func (s *Task) SetStatus(taskID, status string) error {
	return s.repo.SetStatus(taskID, status)
}

func (s *Task) GetResult(taskID string) (string, error) {
	return s.repo.GetResult(taskID)
}

func (s *Task) SetResult(taskID string, result string) error {
	return s.repo.SetResult(taskID, result)
}
