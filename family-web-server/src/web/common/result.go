package common

type Result struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	Data    any    `json:"data,omitempty"`
}

func NewSuccessResult(data any) *Result {
	return &Result{Code: Success, Message: "操作成功", Data: data}
}

func NewResult(code string, message string, data any) *Result {
	return &Result{Code: code, Message: message, Data: data}
}
