package router

import (
	"time"

	"github.com/gsabadini/go-clean-architecture/adapter/logger"
	"github.com/gsabadini/go-clean-architecture/adapter/repository"
	"github.com/gsabadini/go-clean-architecture/adapter/validator"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type gorillaMux struct {
	router     *mux.Router
	middleware *negroni.Negroni
	log        logger.Logger
	db         repository.SQL
	validator  validator.Validator
	port       Port
	ctxTimeout time.Duration
}

func newGorillaMux(
	log logger.Logger,
	db repository.SQL,
	validator validator.Validator,
	port Port,
	t time.Duration,
) *gorillaMux {
	_ = "STUB: not implemented"
	return nil
}

func (g gorillaMux) Listen() { _ = "STUB: not implemented"; return }

func (g gorillaMux) setAppHandlers(router *mux.Router) { _ = "STUB: not implemented"; return }

func (g gorillaMux) buildCreateTransferAction() *negroni.Negroni {
	_ = "STUB: not implemented"
	return nil
}

func (g gorillaMux) buildFindAllTransferAction() *negroni.Negroni {
	_ = "STUB: not implemented"
	return nil
}

func (g gorillaMux) buildCreateAccountAction() *negroni.Negroni {
	_ = "STUB: not implemented"
	return nil
}

func (g gorillaMux) buildFindAllAccountAction() *negroni.Negroni {
	_ = "STUB: not implemented"
	return nil
}

func (g gorillaMux) buildFindBalanceAccountAction() *negroni.Negroni {
	_ = "STUB: not implemented"
	return nil
}
