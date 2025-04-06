package types

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandlerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUserHandlerRequest(r *http.Request) (*UserHandlerRequest, error) {
	var req UserHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}
	if req.Username == "" || req.Password == "" {
		return nil, fmt.Errorf("Username and password are required")
	}
	return &req, nil
}

type SessionResponse struct {
	Token string `json:"token"`
}
