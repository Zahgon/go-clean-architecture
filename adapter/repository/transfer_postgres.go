package repository

import (
	"context"

	"github.com/gsabadini/go-clean-architecture/domain"
)

type TransferSQL struct {
	db SQL
}

func NewTransferSQL(db SQL) TransferSQL { _ = "STUB: not implemented"; return *new(TransferSQL) }

func (t TransferSQL) Create(ctx context.Context, transfer domain.Transfer) (domain.Transfer, error) {
	_ = "STUB: not implemented"
	return *new(domain.Transfer), nil
}

func (t TransferSQL) FindAll(ctx context.Context) ([]domain.Transfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t TransferSQL) WithTransaction(ctx context.Context, fn func(ctxTx context.Context) error) error {
	_ = "STUB: not implemented"
	return nil
}
