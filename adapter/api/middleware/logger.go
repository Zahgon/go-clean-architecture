package middleware

import (
	"net/http"

	"github.com/gsabadini/go-clean-architecture/adapter/logger"
)

type Logger struct {
	log logger.Logger
}

func NewLogger(log logger.Logger) Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (l Logger) Execute(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

func getRequestPayload(r *http.Request) (string, error) { _ = "STUB: not implemented"; return "", nil }
