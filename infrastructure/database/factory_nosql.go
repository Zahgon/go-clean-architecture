package database

import (
	"errors"

	"github.com/gsabadini/go-clean-architecture/adapter/repository"
)

var (
	errInvalidNoSQLDatabaseInstance = errors.New("invalid nosql db instance")
)

const (
	InstanceMongoDB int = iota
)

func NewDatabaseNoSQLFactory(instance int) (repository.NoSQL, error) {
	_ = "STUB: not implemented"
	return *new(repository.NoSQL), nil
}
