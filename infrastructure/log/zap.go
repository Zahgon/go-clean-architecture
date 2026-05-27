package log

import (
	"github.com/gsabadini/go-clean-architecture/adapter/logger"
	"go.uber.org/zap"
)

type zapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() (logger.Logger, error) {
	_ = "STUB: not implemented"
	return *new(logger.Logger), nil
}

func (l *zapLogger) Infof(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *zapLogger) Warnf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *zapLogger) Errorf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *zapLogger) Fatalln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *zapLogger) WithFields(fields logger.Fields) logger.Logger {
	_ = "STUB: not implemented"
	return *new(logger.Logger)
}

func (l *zapLogger) WithError(err error) logger.Logger {
	_ = "STUB: not implemented"
	return *new(logger.Logger)
}
