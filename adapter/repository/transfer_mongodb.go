package repository

import (
	"context"
	"time"

	"github.com/gsabadini/go-clean-architecture/domain"
)

type transferBSON struct {
	ID                   string    `bson:"id"`
	AccountOriginID      string    `bson:"account_origin_id"`
	AccountDestinationID string    `bson:"account_destination_id"`
	Amount               int64     `bson:"amount"`
	CreatedAt            time.Time `bson:"created_at"`
}

type TransferNoSQL struct {
	collectionName string
	db             NoSQL
}

func NewTransferNoSQL(db NoSQL) TransferNoSQL {
	_ = "STUB: not implemented"
	return *new(TransferNoSQL)
}

func (t TransferNoSQL) Create(ctx context.Context, transfer domain.Transfer) (domain.Transfer, error) {
	_ = "STUB: not implemented"
	return *new(domain.Transfer), nil
}

func (t TransferNoSQL) FindAll(ctx context.Context) ([]domain.Transfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t TransferNoSQL) WithTransaction(ctx context.Context, fn func(ctxTx context.Context) error) error {
	_ = "STUB: not implemented"
	return nil
}
