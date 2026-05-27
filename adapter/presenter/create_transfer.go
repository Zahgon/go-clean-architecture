package presenter

import (
	"github.com/gsabadini/go-clean-architecture/domain"
	"github.com/gsabadini/go-clean-architecture/usecase"
)

type createTransferPresenter struct{}

func NewCreateTransferPresenter() usecase.CreateTransferPresenter {
	_ = "STUB: not implemented"
	return *new(usecase.CreateTransferPresenter)
}

func (c createTransferPresenter) Output(transfer domain.Transfer) usecase.CreateTransferOutput {
	_ = "STUB: not implemented"
	return *new(usecase.CreateTransferOutput)
}
