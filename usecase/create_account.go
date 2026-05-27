package usecase

import (
	"context"
	"time"

	"github.com/gsabadini/go-clean-architecture/domain"
)

type (
	// CreateAccountUseCase input port
	CreateAccountUseCase interface {
		Execute(context.Context, CreateAccountInput) (CreateAccountOutput, error)
	}

	// CreateAccountInput input data
	CreateAccountInput struct {
		Name    string `json:"name" validate:"required"`
		CPF     string `json:"cpf" validate:"required"`
		Balance int64  `json:"balance" validate:"gt=0,required"`
	}

	// CreateAccountPresenter output port
	CreateAccountPresenter interface {
		Output(domain.Account) CreateAccountOutput
	}

	// CreateAccountOutput output data
	CreateAccountOutput struct {
		ID        string  `json:"id"`
		Name      string  `json:"name"`
		CPF       string  `json:"cpf"`
		Balance   float64 `json:"balance"`
		CreatedAt string  `json:"created_at"`
	}

	createAccountInteractor struct {
		repo       domain.AccountRepository
		presenter  CreateAccountPresenter
		ctxTimeout time.Duration
	}
)

// NewCreateAccountInteractor creates new createAccountInteractor with its dependencies
func NewCreateAccountInteractor(
	repo domain.AccountRepository,
	presenter CreateAccountPresenter,
	t time.Duration,
) CreateAccountUseCase {
	_ = "STUB: not implemented"
	return *new(CreateAccountUseCase)
}

// Execute orchestrates the use case
func (a createAccountInteractor) Execute(ctx context.Context, input CreateAccountInput) (CreateAccountOutput, error) {
	_ = "STUB: not implemented"
	return *new(CreateAccountOutput), nil
}
