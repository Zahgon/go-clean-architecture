package response

import (
	"net/http"

	"github.com/pkg/errors"
)

var (
	ErrParameterInvalid = errors.New("parameter invalid")

	ErrInvalidInput = errors.New("invalid input")
)

type Error struct {
	statusCode int
	Errors     []string `json:"errors"`
}

func NewError(err error, status int) *Error { _ = "STUB: not implemented"; return nil }

func NewErrorMessage(messages []string, status int) *Error { _ = "STUB: not implemented"; return nil }

func (e Error) Send(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }
