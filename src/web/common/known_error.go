package common

import "encoding/json"

type KnownError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewKnownError(code string, message string) *KnownError {
	return &KnownError{Code: code, Message: message}
}

func (k *KnownError) Error() string {
	return k.String()
}

func (k *KnownError) String() string {
	marshal, err := json.Marshal(k)
	if err != nil {
		return ""
	}
	return string(marshal)
}
