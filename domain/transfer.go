package domain

import (
	"context"
	"time"
)

type TransferID string

func (t TransferID) String() string { _ = "STUB: not implemented"; return "" }

type (
	TransferRepository interface {
		Create(context.Context, Transfer) (Transfer, error)
		FindAll(context.Context) ([]Transfer, error)
		WithTransaction(context.Context, func(context.Context) error) error
	}

	Transfer struct {
		id                   TransferID
		accountOriginID      AccountID
		accountDestinationID AccountID
		amount               Money
		createdAt            time.Time
	}
)

func NewTransfer(
	ID TransferID,
	accountOriginID AccountID,
	accountDestinationID AccountID,
	amount Money,
	createdAt time.Time,
) Transfer {
	_ = "STUB: not implemented"
	return *new(Transfer)
}

func (t Transfer) ID() TransferID { _ = "STUB: not implemented"; return *new(TransferID) }

func (t Transfer) AccountOriginID() AccountID { _ = "STUB: not implemented"; return *new(AccountID) }

func (t Transfer) AccountDestinationID() AccountID {
	_ = "STUB: not implemented"
	return *new(AccountID)
}

func (t Transfer) Amount() Money { _ = "STUB: not implemented"; return *new(Money) }

func (t Transfer) CreatedAt() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
