package response

import (
	"net/http"
)

type Success struct {
	statusCode int
	result     interface{}
}

func NewSuccess(result interface{}, status int) Success {
	_ = "STUB: not implemented"
	return *new(Success)
}

func (r Success) Send(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }
