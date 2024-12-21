package common

type Result struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewSuccessResult() *Result {
	return &Result{Code: Success, Message: "操作成功"}
}

func NewSuccessResultWithData(data any) *Result {
	return &Result{Code: Success, Message: "操作成功", Data: data}
}

func NewResult(code string, message string, data any) *Result {
	return &Result{Code: code, Message: message, Data: data}
}
