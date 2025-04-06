package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	metrics "net/http"
	"project/pkg/postgres_connect"
	"project/processor/api/rabbitMQ"
	"project/processor/config"
	"project/processor/repository/postgres"
	"project/processor/repository/prometheus"
	"project/processor/usecases"
	"project/processor/usecases/service"
)

func main() {
	appFlags := config.ParseFlags()
	var cfg config.AppConfig
	config.MustLoad(appFlags, &cfg)

	metricsRepo := prometheus.NewPrometheusStorage()
	metricsRepo.Register()
	metrics.Handle("/metrics", promhttp.Handler())
	go func() {
		addr := fmt.Sprintf("0.0.0.0:%d", cfg.PrometheusConfig.Port)
		log.Printf("Starting Prometheus metrics server on %s", addr)
		if err := metrics.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("Prometheus server failed: %v", err)
		}
	}()

	storage, err := postgres_connect.NewPostgresStorage(cfg.Postgres, false)
	if err != nil {
		log.Fatalf("failed creating Postgres: %s", err.Error())
	}
	taskRepo := postgres.NewTaskStorage(storage)

	dockerClient, err := usecases.NewDockerClient()
	if err != nil {
		log.Fatalf("failed to create Docker client: %v", err)
	}

	taskProcessor := service.NewProcessor(dockerClient, taskRepo, metricsRepo)

	taskReceiver, err := rabbitMQ.NewRabbitMQReceiver(cfg.RabbitMQ, taskProcessor)
	if err != nil {
		log.Fatalf("failed to create RabbitMQ receiver: %v", err)
	}
	if err := taskReceiver.Receive(); err != nil {
		log.Fatalf("failed to start consumer: %v", err)
	}
	select {}
}
