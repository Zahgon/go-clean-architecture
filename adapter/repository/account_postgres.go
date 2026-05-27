package repository

import (
	"context"

	"github.com/gsabadini/go-clean-architecture/domain"
)

type AccountSQL struct {
	db SQL
}

func NewAccountSQL(db SQL) AccountSQL { _ = "STUB: not implemented"; return *new(AccountSQL) }

func (a AccountSQL) Create(ctx context.Context, account domain.Account) (domain.Account, error) {
	_ = "STUB: not implemented"
	return *new(domain.Account), nil
}

func (a AccountSQL) UpdateBalance(ctx context.Context, ID domain.AccountID, balance domain.Money) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AccountSQL) FindAll(ctx context.Context) ([]domain.Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a AccountSQL) FindByID(ctx context.Context, ID domain.AccountID) (domain.Account, error) {
	_ = "STUB: not implemented"
	return *new(domain.Account), nil
}

func (a AccountSQL) FindBalance(ctx context.Context, ID domain.AccountID) (domain.Account, error) {
	_ = "STUB: not implemented"
	return *new(domain.Account), nil
}
