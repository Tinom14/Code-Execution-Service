package repository

import "project/http_service/domain"

type UserStorage interface {
	Register(login string, password string) error
	Login(login string) (domain.User, error)
}
