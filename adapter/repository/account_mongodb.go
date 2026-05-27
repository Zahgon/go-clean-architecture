package repository

import (
	"context"
	"time"

	"github.com/gsabadini/go-clean-architecture/domain"
)

type accountBSON struct {
	ID        string    `bson:"id"`
	Name      string    `bson:"name"`
	CPF       string    `bson:"cpf"`
	Balance   int64     `bson:"balance"`
	CreatedAt time.Time `bson:"created_at"`
}

type AccountNoSQL struct {
	collectionName string
	db             NoSQL
}

func NewAccountNoSQL(db NoSQL) AccountNoSQL { _ = "STUB: not implemented"; return *new(AccountNoSQL) }

func (a AccountNoSQL) Create(ctx context.Context, account domain.Account) (domain.Account, error) {
	_ = "STUB: not implemented"
	return *new(domain.Account), nil
}

func (a AccountNoSQL) UpdateBalance(ctx context.Context, ID domain.AccountID, balance domain.Money) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AccountNoSQL) FindAll(ctx context.Context) ([]domain.Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a AccountNoSQL) FindByID(ctx context.Context, ID domain.AccountID) (domain.Account, error) {
	_ = "STUB: not implemented"
	return *new(domain.Account), nil
}

func (a AccountNoSQL) FindBalance(ctx context.Context, ID domain.AccountID) (domain.Account, error) {
	_ = "STUB: not implemented"
	return *new(domain.Account), nil
}
