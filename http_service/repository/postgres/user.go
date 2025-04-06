package postgres

import (
	"fmt"
	_ "github.com/lib/pq"
	"project/http_service/domain"
	"project/http_service/repository"
	"project/pkg/postgres_connect"
)

type UserStorage struct {
	users *postgres_connect.PostgresStorage
}

func NewUserStorage(users *postgres_connect.PostgresStorage) *UserStorage {
	return &UserStorage{users: users}
}

func (us *UserStorage) Register(login string, password string) error {
	_, err := us.users.Db.Exec("INSERT INTO Users (login, password) VALUES ($1, $2)", login, password)
	if err != nil {
		return fmt.Errorf("failed to register user: %v", err)
	}
	return nil
}

func (us *UserStorage) Login(login string) (domain.User, error) {
	row := us.users.Db.QueryRow("SELECT * FROM Users WHERE login = $1", login)
	var user domain.User
	err := row.Scan(&user.UserID, &user.Login, &user.Password)
	if err != nil {
		return domain.User{}, repository.NotFound
	}
	return user, nil
}
