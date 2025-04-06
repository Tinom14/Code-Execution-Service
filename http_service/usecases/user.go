package usecases

type UserService interface {
	Register(login string, password string) error
	Login(login string, password string) (int, error)
}
