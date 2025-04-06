package usecases

type SessionService interface {
	CreateSession(userID int) (string, error)
	CheckSession(sessionID string) error
}
