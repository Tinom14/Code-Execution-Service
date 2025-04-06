package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"project/http_service/repository"
)

type Session struct {
	repo repository.SessionStorage
}

func NewSessionService(repo repository.SessionStorage) *Session {
	return &Session{
		repo: repo,
	}
}

func (s *Session) CreateSession(userID int) (string, error) {
	var sessionID = GeneratorSessionID()
	return s.repo.CreateSession(sessionID, userID)
}

func (s *Session) CheckSession(sessionID string) error {
	err := s.repo.CheckSession(sessionID)
	if err != nil {
		return errors.New("wrong sessionID")
	}
	return nil
}

func GeneratorSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
