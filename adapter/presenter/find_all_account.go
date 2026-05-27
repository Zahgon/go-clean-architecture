package presenter

import (
	"github.com/gsabadini/go-clean-architecture/domain"
	"github.com/gsabadini/go-clean-architecture/usecase"
)

type findAllAccountPresenter struct{}

func NewFindAllAccountPresenter() usecase.FindAllAccountPresenter {
	_ = "STUB: not implemented"
	return *new(usecase.FindAllAccountPresenter)
}

func (a findAllAccountPresenter) Output(accounts []domain.Account) []usecase.FindAllAccountOutput {
	_ = "STUB: not implemented"
	return nil
}
