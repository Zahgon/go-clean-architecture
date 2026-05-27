package log

import (
	"github.com/gsabadini/go-clean-architecture/adapter/logger"
	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger() logger.Logger { _ = "STUB: not implemented"; return *new(logger.Logger) }

func (l *logrusLogger) Infof(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logrusLogger) Warnf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *logrusLogger) Fatalln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logrusLogger) WithFields(fields logger.Fields) logger.Logger {
	_ = "STUB: not implemented"
	return *new(logger.Logger)
}

func (l *logrusLogger) WithError(err error) logger.Logger {
	_ = "STUB: not implemented"
	return *new(logger.Logger)
}

type logrusLogEntry struct {
	entry *logrus.Entry
}

func (l *logrusLogEntry) Infof(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *logrusLogEntry) Warnf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *logrusLogEntry) Errorf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *logrusLogEntry) Fatalln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logrusLogEntry) WithFields(fields logger.Fields) logger.Logger {
	_ = "STUB: not implemented"
	return *new(logger.Logger)
}

func (l *logrusLogEntry) WithError(err error) logger.Logger {
	_ = "STUB: not implemented"
	return *new(logger.Logger)
}

func convertToLogrusFields(fields logger.Fields) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}
