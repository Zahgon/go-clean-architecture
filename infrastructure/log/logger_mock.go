package log

import "github.com/gsabadini/go-clean-architecture/adapter/logger"

type LoggerMock struct{}

func (l LoggerMock) Infof(_ string, _ ...interface{})  { _ = "STUB: not implemented"; return }
func (l LoggerMock) Warnf(_ string, _ ...interface{})  { _ = "STUB: not implemented"; return }
func (l LoggerMock) Errorf(_ string, _ ...interface{}) { _ = "STUB: not implemented"; return }
func (l LoggerMock) Fatalln(_ ...interface{})          { _ = "STUB: not implemented"; return }
func (l LoggerMock) WithFields(_ logger.Fields) logger.Logger {
	_ = "STUB: not implemented"
	return *new(logger.Logger)
}
func (l LoggerMock) WithError(_ error) logger.Logger {
	_ = "STUB: not implemented"
	return *new(logger.Logger)
}

type LoggerEntryMock struct{}

func (l LoggerEntryMock) Infof(_ string, _ ...interface{})  { _ = "STUB: not implemented"; return }
func (l LoggerEntryMock) Warnf(_ string, _ ...interface{})  { _ = "STUB: not implemented"; return }
func (l LoggerEntryMock) Errorf(_ string, _ ...interface{}) { _ = "STUB: not implemented"; return }
func (l LoggerEntryMock) Fatalln(_ ...interface{})          { _ = "STUB: not implemented"; return }
func (l LoggerEntryMock) WithFields(_ logger.Fields) logger.Logger {
	_ = "STUB: not implemented"
	return *new(logger.Logger)
}
func (l LoggerEntryMock) WithError(_ error) logger.Logger {
	_ = "STUB: not implemented"
	return *new(logger.Logger)
}
