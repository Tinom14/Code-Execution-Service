package domain

type Session struct {
	UserID    int    `json:"user_id"`
	SessionID string `json:"session_id"`
}
