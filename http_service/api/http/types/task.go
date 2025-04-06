package types

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type PostTaskHandlerRequest struct {
	Translator string `json:"translator"`
	Code       string `json:"code"`
}

func CreatePostTaskHandlerRequest(r *http.Request) (*PostTaskHandlerRequest, error) {
	var req PostTaskHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}
	return &req, nil
}

type PostTaskHandlerResponse struct {
	TaskID string `json:"task_id"`
}

type GetTaskStatusHandlerRequest struct {
	TaskID string
}

func CreateGetTaskStatusHandlerRequest(r *http.Request) *GetTaskStatusHandlerRequest {
	taskID := chi.URLParam(r, "task_id")
	return &GetTaskStatusHandlerRequest{TaskID: taskID}
}

type GetTaskStatusHandlerResponse struct {
	Status string `json:"status"`
}

type GetTaskResultHandlerRequest struct {
	TaskID string
}

func CreateGetTaskResultHandlerRequest(r *http.Request) *GetTaskResultHandlerRequest {
	taskID := chi.URLParam(r, "task_id")
	return &GetTaskResultHandlerRequest{TaskID: taskID}
}

type GetTaskResultHandlerResponse struct {
	Result string `json:"result"`
}

type PostTaskCommitRequest struct {
	TaskID string `json:"task_id"`
	Result string `json:"result"`
}

type PostTaskCommitResponse struct {
	TaskID string `json:"task_id"`
	Result string `json:"result"`
}
