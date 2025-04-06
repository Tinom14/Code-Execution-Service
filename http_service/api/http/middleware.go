package http

import (
	"net/http"
	"project/http_service/usecases"
	"strings"
)

// MiddlewareAuth проверяет, есть ли у пользователя действительный токен.
// @Summary Проверка токена
// @Description Позволяет доступ только авторизованным пользователям
// @Tags auth
// @Security ApiKeyAuth
// @Param Authorization header string true "Токен авторизации, формат: Bearer {token}"
// @Failure 401  {object} string
// @Router /protected-endpoint [get]
func MiddlewareAuth(sessionService usecases.SessionService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			sessionID := strings.TrimPrefix(authHeader, "Bearer ")
			err := sessionService.CheckSession(sessionID)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
