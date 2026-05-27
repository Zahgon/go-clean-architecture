package logging

import (
	"github.com/gsabadini/go-clean-architecture/adapter/logger"
)

type Info struct {
	log        logger.Logger
	key        string
	httpStatus int
}

func NewInfo(log logger.Logger, key string, httpStatus int) Info {
	_ = "STUB: not implemented"
	return *new(Info)
}

func (i Info) Log(msg string) { _ = "STUB: not implemented"; return }
