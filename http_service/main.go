package main

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"project/http_service/api/http"
	"project/http_service/config"
	_ "project/http_service/docs"
	pkgHttp "project/http_service/pkg/http"
	"project/http_service/repository/postgres"
	rabbitMQ "project/http_service/repository/rabbit_mq"
	"project/http_service/repository/redis"
	"project/http_service/usecases/service"
	"project/pkg/postgres_connect"
	"time"
)

// @title Task API
// @version 1.0
// @description API для создания задач и получения их статуса/результата.
// @host localhost:8080
// @BasePath /
func main() {
	appFlags := config.ParseFlags()
	var cfg config.AppConfig
	config.MustLoad(appFlags.ConfigPath, &cfg)

	storage, err := postgres_connect.NewPostgresStorage(cfg.Postgres, cfg.MigrationsEnabled)
	if err != nil {
		log.Fatalf("failed creating Postgres: %s", err.Error())
	}

	taskRepo := postgres.NewTaskStorage(storage)
	taskSender, err := rabbitMQ.NewRabbitMQSender(cfg.RabbitMQ)
	if err != nil {
		log.Fatalf("failed creating rabbitMQ: %s", err.Error())
	}
	taskService := service.NewTaskService(taskRepo, taskSender)
	taskHandlers := http.NewTaskServer(taskService)

	userRepo := postgres.NewUserStorage(storage)
	userService := service.NewUserService(userRepo)

	sessionRepo, err := redis.NewRedisStorage(cfg.Redis, 24*time.Hour)
	if err != nil {
		log.Fatalf("failed creating redis: %s", err.Error())
	}
	sessionService := service.NewSessionService(sessionRepo)

	authHandlers := http.NewAuthServer(userService, sessionService)

	r := chi.NewRouter()
	authHandlers.WithAuthHandlers(r)

	protectedRoutes := chi.NewRouter()
	protectedRoutes.Use(http.MiddlewareAuth(sessionService))
	taskHandlers.WithTaskHandlers(protectedRoutes)
	r.Mount("/", protectedRoutes)
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	log.Printf("Starting server on %s", cfg.Address)
	if err := pkgHttp.CreateAndRunServer(r, cfg.Address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
