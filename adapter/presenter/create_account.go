package presenter

import (
	"github.com/gsabadini/go-clean-architecture/domain"
	"github.com/gsabadini/go-clean-architecture/usecase"
)

type createAccountPresenter struct{}

func NewCreateAccountPresenter() usecase.CreateAccountPresenter {
	_ = "STUB: not implemented"
	return *new(usecase.CreateAccountPresenter)
}

func (a createAccountPresenter) Output(account domain.Account) usecase.CreateAccountOutput {
	_ = "STUB: not implemented"
	return *new(usecase.CreateAccountOutput)
}
