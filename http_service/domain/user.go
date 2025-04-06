package domain

type User struct {
	UserID   int    `json:"id"`
	Password string `json:"password"`
	Login    string `json:"login"`
}
