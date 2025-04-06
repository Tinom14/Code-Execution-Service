package http

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"project/http_service/api/http/types"
	"project/http_service/usecases"
)

type TaskServer struct {
	service usecases.TaskService
}

func NewTaskServer(storage usecases.TaskService) *TaskServer {
	return &TaskServer{service: storage}
}

// PostTaskHandler создает новую задачу
// @Summary Создать задачу
// @Description Запускает новую задачу и возвращает ее ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Токен авторизации, формат: Bearer {token}"
// @Success 201 {object} types.PostTaskHandlerResponse
// @Failure 401 {object} string
// @Failure 500 {object} string
// @Router /task [post]
func (s *TaskServer) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostTaskHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	taskID, err := s.service.CreateTask(req.Code, req.Translator)
	if err != nil {
		types.ProcessError(w, err, nil)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&types.PostTaskHandlerResponse{TaskID: taskID})
}

// GetStatusHandler получает статус задачи
// @Summary Получить статус задачи
// @Description Возвращает текущий статус задачи по ID
// @Tags tasks
// @Param task_id path string true "ID задачи"
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Токен авторизации, формат: Bearer {token}"
// @Success 200 {object} types.GetTaskStatusHandlerResponse
// @Failure 401 {object} string
// @Failure 404 {object} string
// @Router /status/{task_id} [get]
func (s *TaskServer) GetStatusHandler(w http.ResponseWriter, r *http.Request) {
	req := types.CreateGetTaskStatusHandlerRequest(r)
	status, err := s.service.GetStatus(req.TaskID)
	types.ProcessError(w, err, &types.GetTaskStatusHandlerResponse{Status: status})
}

// GetResultHandler получает результат задачи
// @Summary Получить результат задачи
// @Description Возвращает результат выполнения задачи по ID
// @Tags tasks
// @Param task_id path string true "ID задачи"
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Токен авторизации, формат: Bearer {token}"
// @Success 200 {object} types.GetTaskResultHandlerResponse
// @Failure 401 {object} string
// @Failure 404 {object} string
// @Router /result/{task_id} [get]
func (s *TaskServer) GetResultHandler(w http.ResponseWriter, r *http.Request) {
	req := types.CreateGetTaskResultHandlerRequest(r)
	result, err := s.service.GetResult(req.TaskID)
	types.ProcessError(w, err, &types.GetTaskResultHandlerResponse{Result: result})
}

func (s *TaskServer) WithTaskHandlers(r chi.Router) {
	r.Post("/task", s.PostTaskHandler)
	r.Get("/status/{task_id}", s.GetStatusHandler)
	r.Get("/result/{task_id}", s.GetResultHandler)
}
