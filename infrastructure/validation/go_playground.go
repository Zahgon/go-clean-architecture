package validation

import (
	"github.com/gsabadini/go-clean-architecture/adapter/validator"

	ut "github.com/go-playground/universal-translator"
	go_playground "github.com/go-playground/validator/v10"
)

type goPlayground struct {
	validator *go_playground.Validate
	translate ut.Translator
	err       error
	msg       []string
}

func NewGoPlayground() (validator.Validator, error) {
	_ = "STUB: not implemented"
	return *new(validator.Validator), nil
}

func (g *goPlayground) Validate(i interface{}) error { _ = "STUB: not implemented"; return nil }

func (g *goPlayground) Messages() []string { _ = "STUB: not implemented"; return nil }
