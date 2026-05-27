package presenter

import (
	"github.com/gsabadini/go-clean-architecture/domain"
	"github.com/gsabadini/go-clean-architecture/usecase"
)

type findAccountBalancePresenter struct{}

func NewFindAccountBalancePresenter() usecase.FindAccountBalancePresenter {
	_ = "STUB: not implemented"
	return *new(usecase.FindAccountBalancePresenter)
}

func (a findAccountBalancePresenter) Output(balance domain.Money) usecase.FindAccountBalanceOutput {
	_ = "STUB: not implemented"
	return *new(usecase.FindAccountBalanceOutput)
}
