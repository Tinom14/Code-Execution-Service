package ram_storage

import (
	"project/http_service/domain"
	"project/http_service/repository"
	"sync"
)

type SessionRamStorage struct {
	mu       sync.Mutex
	sessions map[string]domain.Session
}

func NewSessionRamStorage() *SessionRamStorage {
	return &SessionRamStorage{
		sessions: make(map[string]domain.Session),
	}
}

func (s *SessionRamStorage) CreateSession(sessionID string, userID int) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sessions[sessionID] = domain.Session{UserID: userID}
	return sessionID, nil
}

func (s *SessionRamStorage) CheckSession(sessionID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.sessions[sessionID]
	if !exists {
		return repository.NotFound
	}
	return nil
}
