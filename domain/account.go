package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrAccountNotFound = errors.New("account not found")

	ErrAccountOriginNotFound = errors.New("account origin not found")

	ErrAccountDestinationNotFound = errors.New("account destination not found")

	ErrInsufficientBalance = errors.New("origin account does not have sufficient balance")
)

type AccountID string

func (a AccountID) String() string { _ = "STUB: not implemented"; return "" }

type (
	AccountRepository interface {
		Create(context.Context, Account) (Account, error)
		UpdateBalance(context.Context, AccountID, Money) error
		FindAll(context.Context) ([]Account, error)
		FindByID(context.Context, AccountID) (Account, error)
		FindBalance(context.Context, AccountID) (Account, error)
	}

	Account struct {
		id        AccountID
		name      string
		cpf       string
		balance   Money
		createdAt time.Time
	}
)

func NewAccount(ID AccountID, name, CPF string, balance Money, createdAt time.Time) Account {
	_ = "STUB: not implemented"
	return *new(Account)
}

func (a *Account) Deposit(amount Money) { _ = "STUB: not implemented"; return }

func (a *Account) Withdraw(amount Money) error { _ = "STUB: not implemented"; return nil }

func (a Account) ID() AccountID { _ = "STUB: not implemented"; return *new(AccountID) }

func (a Account) Name() string { _ = "STUB: not implemented"; return "" }

func (a Account) CPF() string { _ = "STUB: not implemented"; return "" }

func (a Account) Balance() Money { _ = "STUB: not implemented"; return *new(Money) }

func (a Account) CreatedAt() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func NewAccountBalance(balance Money) Account { _ = "STUB: not implemented"; return *new(Account) }
