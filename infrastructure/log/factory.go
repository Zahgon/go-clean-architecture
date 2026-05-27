package log

import (
	"errors"

	"github.com/gsabadini/go-clean-architecture/adapter/logger"
)

const (
	InstanceZapLogger int = iota
	InstanceLogrusLogger
)

var (
	errInvalidLoggerInstance = errors.New("invalid log instance")
)

func NewLoggerFactory(instance int) (logger.Logger, error) {
	_ = "STUB: not implemented"
	return *new(logger.Logger), nil
}
