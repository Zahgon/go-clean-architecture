package validation

import (
	"errors"

	"github.com/gsabadini/go-clean-architecture/adapter/validator"
)

var (
	errInvalidValidatorInstance = errors.New("invalid validator instance")
)

const (
	InstanceGoPlayground int = iota
)

func NewValidatorFactory(instance int) (validator.Validator, error) {
	_ = "STUB: not implemented"
	return *new(validator.Validator), nil
}
