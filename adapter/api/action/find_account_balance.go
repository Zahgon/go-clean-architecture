package action

import (
	"net/http"

	"github.com/gsabadini/go-clean-architecture/adapter/logger"
	"github.com/gsabadini/go-clean-architecture/usecase"
)

type FindAccountBalanceAction struct {
	uc  usecase.FindAccountBalanceUseCase
	log logger.Logger
}

func NewFindAccountBalanceAction(uc usecase.FindAccountBalanceUseCase, log logger.Logger) FindAccountBalanceAction {
	_ = "STUB: not implemented"
	return *new(FindAccountBalanceAction)
}

func (a FindAccountBalanceAction) Execute(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
