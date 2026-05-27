package database

import (
	"time"
)

type config struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string

	ctxTimeout time.Duration
}

func newConfigMongoDB() *config { _ = "STUB: not implemented"; return nil }

func newConfigPostgres() *config { _ = "STUB: not implemented"; return nil }
