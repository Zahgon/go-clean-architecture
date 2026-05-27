package action

import (
	"net/http"

	"github.com/gsabadini/go-clean-architecture/adapter/logger"
	"github.com/gsabadini/go-clean-architecture/usecase"
)

type FindAllTransferAction struct {
	uc  usecase.FindAllTransferUseCase
	log logger.Logger
}

func NewFindAllTransferAction(uc usecase.FindAllTransferUseCase, log logger.Logger) FindAllTransferAction {
	_ = "STUB: not implemented"
	return *new(FindAllTransferAction)
}

func (t FindAllTransferAction) Execute(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
