package logging

import (
	"github.com/gsabadini/go-clean-architecture/adapter/logger"
)

type Error struct {
	log        logger.Logger
	err        error
	key        string
	httpStatus int
}

func NewError(log logger.Logger, err error, key string, httpStatus int) Error {
	_ = "STUB: not implemented"
	return *new(Error)
}

func (e Error) Log(msg string) { _ = "STUB: not implemented"; return }
