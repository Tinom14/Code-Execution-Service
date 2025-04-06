package ram_storage

import (
	"project/http_service/domain"
	"project/http_service/repository"
	"sync"
)

type UserRamStorage struct {
	mu    sync.Mutex
	users map[string]domain.User
}

func NewUserRamStorage() *UserRamStorage {
	return &UserRamStorage{
		users: make(map[string]domain.User),
	}
}

func (u *UserRamStorage) Register(login string, password string) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.users[login] = domain.User{UserID: len(u.users), Password: password}
	return nil
}

func (u *UserRamStorage) Login(login string) (domain.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	user, exists := u.users[login]
	if !exists {
		return domain.User{}, repository.NotFound
	}
	return user, nil
}
