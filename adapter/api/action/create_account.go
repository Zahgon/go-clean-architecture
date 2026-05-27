package action

import (
	"net/http"

	"github.com/gsabadini/go-clean-architecture/adapter/logger"
	"github.com/gsabadini/go-clean-architecture/adapter/validator"
	"github.com/gsabadini/go-clean-architecture/usecase"
)

type CreateAccountAction struct {
	uc        usecase.CreateAccountUseCase
	log       logger.Logger
	validator validator.Validator
}

func NewCreateAccountAction(uc usecase.CreateAccountUseCase, log logger.Logger, v validator.Validator) CreateAccountAction {
	_ = "STUB: not implemented"
	return *new(CreateAccountAction)
}

func (a CreateAccountAction) Execute(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (a CreateAccountAction) validateInput(input usecase.CreateAccountInput) []string {
	_ = "STUB: not implemented"
	return nil
}

func (a CreateAccountAction) cleanCPF(cpf string) string { _ = "STUB: not implemented"; return "" }
