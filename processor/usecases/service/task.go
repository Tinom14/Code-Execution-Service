package service

import (
	"log"
	"project/http_service/api/http/types"
	"project/processor/repository"
	"project/processor/repository/prometheus"
	"project/processor/usecases"
	"strings"
	"time"
)

type Processor struct {
	dockerCli   *usecases.DockerClient
	taskRepo    repository.TaskSender
	metricsRepo *prometheus.PrometheusStorage
}

func NewProcessor(dockerCli *usecases.DockerClient, taskRepo repository.TaskSender, metricsRepo *prometheus.PrometheusStorage) *Processor {
	return &Processor{
		dockerCli:   dockerCli,
		taskRepo:    taskRepo,
		metricsRepo: metricsRepo,
	}
}

func (p *Processor) CompleteTask(taskID string, language string, code string) error {
	startTime := time.Now()
	defer func() {
		p.metricsRepo.RecordTaskDuration(language, time.Since(startTime))
		p.metricsRepo.RecordLanguageUsage(language)
	}()

	result, err := p.dockerCli.RunCodeInContainer(language, code)
	if err != nil {
		log.Printf("error running code: %v", err)
	}

	log.Printf("TaskID: %s result: %s", taskID, result)
	cleanResult := strings.ReplaceAll(result, "\n", "")
	payload := types.PostTaskCommitRequest{
		TaskID: taskID,
		Result: cleanResult,
	}

	if err := p.taskRepo.SendResult(payload); err != nil {
		log.Printf("failed to send result to server: %v", err)
	}
	return nil
}
