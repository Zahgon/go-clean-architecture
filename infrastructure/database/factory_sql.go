package database

import (
	"errors"

	"github.com/gsabadini/go-clean-architecture/adapter/repository"
)

var (
	errInvalidSQLDatabaseInstance = errors.New("invalid sql db instance")
)

const (
	InstancePostgres int = iota
)

func NewDatabaseSQLFactory(instance int) (repository.SQL, error) {
	_ = "STUB: not implemented"
	return *new(repository.SQL), nil
}
