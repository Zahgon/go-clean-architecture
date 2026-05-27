package action

import (
	"net/http"

	"github.com/gsabadini/go-clean-architecture/adapter/logger"
	"github.com/gsabadini/go-clean-architecture/usecase"
)

type FindAllAccountAction struct {
	uc  usecase.FindAllAccountUseCase
	log logger.Logger
}

func NewFindAllAccountAction(uc usecase.FindAllAccountUseCase, log logger.Logger) FindAllAccountAction {
	_ = "STUB: not implemented"
	return *new(FindAllAccountAction)
}

func (a FindAllAccountAction) Execute(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
