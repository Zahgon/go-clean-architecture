package usecase

import (
	"context"
	"time"

	"github.com/gsabadini/go-clean-architecture/domain"
)

type (
	// CreateTransferUseCase input port
	CreateTransferUseCase interface {
		Execute(context.Context, CreateTransferInput) (CreateTransferOutput, error)
	}

	// CreateTransferInput input data
	CreateTransferInput struct {
		AccountOriginID      string `json:"account_origin_id" validate:"required,uuid4"`
		AccountDestinationID string `json:"account_destination_id" validate:"required,uuid4"`
		Amount               int64  `json:"amount" validate:"gt=0,required"`
	}

	// CreateTransferPresenter output port
	CreateTransferPresenter interface {
		Output(domain.Transfer) CreateTransferOutput
	}

	// CreateTransferOutput output data
	CreateTransferOutput struct {
		ID                   string  `json:"id"`
		AccountOriginID      string  `json:"account_origin_id"`
		AccountDestinationID string  `json:"account_destination_id"`
		Amount               float64 `json:"amount"`
		CreatedAt            string  `json:"created_at"`
	}

	createTransferInteractor struct {
		transferRepo domain.TransferRepository
		accountRepo  domain.AccountRepository
		presenter    CreateTransferPresenter
		ctxTimeout   time.Duration
	}
)

// NewCreateTransferInteractor creates new createTransferInteractor with its dependencies
func NewCreateTransferInteractor(
	transferRepo domain.TransferRepository,
	accountRepo domain.AccountRepository,
	presenter CreateTransferPresenter,
	t time.Duration,
) CreateTransferUseCase {
	_ = "STUB: not implemented"
	return *new(CreateTransferUseCase)
}

// Execute orchestrates the use case
func (t createTransferInteractor) Execute(ctx context.Context, input CreateTransferInput) (CreateTransferOutput, error) {
	_ = "STUB: not implemented"
	return *new(CreateTransferOutput), nil
}

func (t createTransferInteractor) process(ctx context.Context, input CreateTransferInput) error {
	_ = "STUB: not implemented"
	return nil
}
