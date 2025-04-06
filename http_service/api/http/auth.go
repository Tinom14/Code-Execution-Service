package http

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"project/http_service/api/http/types"
	"project/http_service/usecases"
)

type AuthServer struct {
	userService    usecases.UserService
	sessionService usecases.SessionService
}

func NewAuthServer(userStorage usecases.UserService, sessionStorage usecases.SessionService) *AuthServer {
	return &AuthServer{userService: userStorage, sessionService: sessionStorage}
}

// RegisterHandler регистрирует нового пользователя
// @Summary      Регистрация пользователя
// @Description  Создаёт нового пользователя
// @Tags register
// @Accept json
// @Produce json
// @Param request body types.UserHandlerRequest true "Данные пользователя"
// @Success 201
// @Failure 400  {object} string
// @Router /register [post]
func (s *AuthServer) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateUserHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = s.userService.Register(req.Username, req.Password)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// LoginHandler аутентифицирует пользователя и выдаёт токен
// @Summary Аутентификация пользователя
// @Description  Проверяет логин и пароль, возвращает токен
// @Tags auth
// @Accept json
// @Produce json
// @Param request body types.UserHandlerRequest true "Данные пользователя"
// @Success 200 {object} types.SessionResponse "Успешный вход"
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Router /login [post]
func (s *AuthServer) LoginHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateUserHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var userID int
	userID, err = s.userService.Login(req.Username, req.Password)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	token, err := s.sessionService.CreateSession(userID)

	json.NewEncoder(w).Encode(types.SessionResponse{Token: token})
}

func (s *AuthServer) WithAuthHandlers(r chi.Router) {
	r.Post("/register", s.RegisterHandler)
	r.Post("/login", s.LoginHandler)
}
