package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"project/http_service/repository"
)

type User struct {
	repo repository.UserStorage
}

func NewUserService(repo repository.UserStorage) *User {
	return &User{
		repo: repo,
	}
}

func (u *User) Register(login string, password string) error {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashPasswordStr := string(hashPassword)
	return u.repo.Register(login, hashPasswordStr)
}

func (u *User) Login(login string, password string) (int, error) {
	user, err := u.repo.Login(login)
	if err != nil {
		return -1, errors.New("wrong login")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return -1, errors.New("wrong password")
	}
	return user.UserID, nil
}
