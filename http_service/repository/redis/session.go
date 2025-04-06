package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"project/http_service/config"
	"project/http_service/repository"
	"strconv"
	"time"
)

type SessionRedisStorage struct {
	client    *redis.Client
	expiresIn time.Duration
}

func NewRedisStorage(cfg config.Redis, expiresIn time.Duration) (*SessionRedisStorage, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + strconv.Itoa(int(cfg.Port)),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &SessionRedisStorage{client: client, expiresIn: expiresIn}, nil
}

func (s *SessionRedisStorage) CreateSession(sessionID string, userID int) (string, error) {
	if err := s.client.Set(context.Background(), sessionID, userID, s.expiresIn).Err(); err != nil {
		return "", fmt.Errorf("failed to save session to Redis: %w", err)
	}
	return sessionID, nil
}

func (s *SessionRedisStorage) CheckSession(sessionID string) error {
	exists, err := s.client.Exists(context.Background(), sessionID).Result()
	if err != nil {
		return fmt.Errorf("failed to check session: %w", err)
	}

	if exists == 0 {
		return repository.NotFound
	}
	return nil
}
