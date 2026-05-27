package action

import (
	"net/http"

	"github.com/gsabadini/go-clean-architecture/adapter/logger"
	"github.com/gsabadini/go-clean-architecture/adapter/validator"
	"github.com/gsabadini/go-clean-architecture/usecase"
)

type CreateTransferAction struct {
	log       logger.Logger
	uc        usecase.CreateTransferUseCase
	validator validator.Validator

	logKey, logMsg string
}

func NewCreateTransferAction(uc usecase.CreateTransferUseCase, log logger.Logger, v validator.Validator) CreateTransferAction {
	_ = "STUB: not implemented"
	return *new(CreateTransferAction)
}

func (t CreateTransferAction) Execute(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (t CreateTransferAction) handleErr(w http.ResponseWriter, err error) {
	_ = "STUB: not implemented"
	return
}

func (t CreateTransferAction) validateInput(input usecase.CreateTransferInput) []string {
	_ = "STUB: not implemented"
	return nil
}
