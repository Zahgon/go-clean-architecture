package database

import (
	"context"
	"database/sql"

	"github.com/gsabadini/go-clean-architecture/adapter/repository"

	_ "github.com/lib/pq"
)

type postgresHandler struct {
	db *sql.DB
}

func NewPostgresHandler(c *config) (*postgresHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p postgresHandler) BeginTx(ctx context.Context) (repository.Tx, error) {
	_ = "STUB: not implemented"
	return *new(repository.Tx), nil
}

func (p postgresHandler) ExecuteContext(ctx context.Context, query string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (p postgresHandler) QueryContext(ctx context.Context, query string, args ...interface{}) (repository.Rows, error) {
	_ = "STUB: not implemented"
	return *new(repository.Rows), nil
}

func (p postgresHandler) QueryRowContext(ctx context.Context, query string, args ...interface{}) repository.Row {
	_ = "STUB: not implemented"
	return *new(repository.Row)
}

type postgresRow struct {
	row *sql.Row
}

func newPostgresRow(row *sql.Row) postgresRow { _ = "STUB: not implemented"; return *new(postgresRow) }

func (pr postgresRow) Scan(dest ...interface{}) error { _ = "STUB: not implemented"; return nil }

type postgresRows struct {
	rows *sql.Rows
}

func newPostgresRows(rows *sql.Rows) postgresRows {
	_ = "STUB: not implemented"
	return *new(postgresRows)
}

func (pr postgresRows) Scan(dest ...interface{}) error { _ = "STUB: not implemented"; return nil }

func (pr postgresRows) Next() bool { _ = "STUB: not implemented"; return false }

func (pr postgresRows) Err() error { _ = "STUB: not implemented"; return nil }

func (pr postgresRows) Close() error { _ = "STUB: not implemented"; return nil }

type postgresTx struct {
	tx *sql.Tx
}

func newPostgresTx(tx *sql.Tx) postgresTx { _ = "STUB: not implemented"; return *new(postgresTx) }

func (p postgresTx) ExecuteContext(ctx context.Context, query string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (p postgresTx) QueryContext(ctx context.Context, query string, args ...interface{}) (repository.Rows, error) {
	_ = "STUB: not implemented"
	return *new(repository.Rows), nil
}

func (p postgresTx) QueryRowContext(ctx context.Context, query string, args ...interface{}) repository.Row {
	_ = "STUB: not implemented"
	return *new(repository.Row)
}

func (p postgresTx) Commit() error { _ = "STUB: not implemented"; return nil }

func (p postgresTx) Rollback() error { _ = "STUB: not implemented"; return nil }
