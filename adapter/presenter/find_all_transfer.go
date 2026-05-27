package presenter

import (
	"github.com/gsabadini/go-clean-architecture/domain"
	"github.com/gsabadini/go-clean-architecture/usecase"
)

type findAllTransferPresenter struct{}

func NewFindAllTransferPresenter() usecase.FindAllTransferPresenter {
	_ = "STUB: not implemented"
	return *new(usecase.FindAllTransferPresenter)
}

func (a findAllTransferPresenter) Output(transfers []domain.Transfer) []usecase.FindAllTransferOutput {
	_ = "STUB: not implemented"
	return nil
}
