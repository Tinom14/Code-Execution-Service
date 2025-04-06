package types

type TaskFromRabbit struct {
	TaskID   string `json:"task_id"`
	Code     string `json:"code"`
	Language string `json:"language"`
}
