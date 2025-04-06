package repository

type SessionStorage interface {
	CreateSession(sessionID string, userID int) (string, error)
	CheckSession(sessionID string) error
}
