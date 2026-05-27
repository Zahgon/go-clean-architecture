package infrastructure

import (
	"time"

	"github.com/gsabadini/go-clean-architecture/adapter/logger"
	"github.com/gsabadini/go-clean-architecture/adapter/repository"
	"github.com/gsabadini/go-clean-architecture/adapter/validator"
	"github.com/gsabadini/go-clean-architecture/infrastructure/router"
)

type config struct {
	appName       string
	logger        logger.Logger
	validator     validator.Validator
	dbSQL         repository.SQL
	dbNoSQL       repository.NoSQL
	ctxTimeout    time.Duration
	webServerPort router.Port
	webServer     router.Server
}

func NewConfig() *config { _ = "STUB: not implemented"; return nil }

func (c *config) ContextTimeout(t time.Duration) *config { _ = "STUB: not implemented"; return nil }

func (c *config) Name(name string) *config { _ = "STUB: not implemented"; return nil }

func (c *config) Logger(instance int) *config { _ = "STUB: not implemented"; return nil }

func (c *config) DbSQL(instance int) *config { _ = "STUB: not implemented"; return nil }

func (c *config) DbNoSQL(instance int) *config { _ = "STUB: not implemented"; return nil }

func (c *config) Validator(instance int) *config { _ = "STUB: not implemented"; return nil }

func (c *config) WebServer(instance int) *config { _ = "STUB: not implemented"; return nil }

func (c *config) WebServerPort(port string) *config { _ = "STUB: not implemented"; return nil }

func (c *config) Start() { _ = "STUB: not implemented"; return }
