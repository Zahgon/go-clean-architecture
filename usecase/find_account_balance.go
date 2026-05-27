package usecase

import (
	"context"
	"time"

	"github.com/gsabadini/go-clean-architecture/domain"
)

type (
	// FindAccountBalanceUseCase input port
	FindAccountBalanceUseCase interface {
		Execute(context.Context, domain.AccountID) (FindAccountBalanceOutput, error)
	}

	// FindAccountBalanceInput input data
	FindAccountBalanceInput struct {
		ID int64 `json:"balance" validate:"gt=0,required"`
	}

	// FindAccountBalancePresenter output port
	FindAccountBalancePresenter interface {
		Output(domain.Money) FindAccountBalanceOutput
	}

	// FindAccountBalanceOutput output data
	FindAccountBalanceOutput struct {
		Balance float64 `json:"balance"`
	}

	findBalanceAccountInteractor struct {
		repo       domain.AccountRepository
		presenter  FindAccountBalancePresenter
		ctxTimeout time.Duration
	}
)

// NewFindBalanceAccountInteractor creates new findBalanceAccountInteractor with its dependencies
func NewFindBalanceAccountInteractor(
	repo domain.AccountRepository,
	presenter FindAccountBalancePresenter,
	t time.Duration,
) FindAccountBalanceUseCase {
	_ = "STUB: not implemented"
	return *new(FindAccountBalanceUseCase)
}

// Execute orchestrates the use case
func (a findBalanceAccountInteractor) Execute(ctx context.Context, ID domain.AccountID) (FindAccountBalanceOutput, error) {
	_ = "STUB: not implemented"
	return *new(FindAccountBalanceOutput), nil
}
