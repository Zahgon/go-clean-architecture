package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gsabadini/go-clean-architecture/adapter/logger"
	"github.com/gsabadini/go-clean-architecture/adapter/repository"
	"github.com/gsabadini/go-clean-architecture/adapter/validator"
)

type ginEngine struct {
	router     *gin.Engine
	log        logger.Logger
	db         repository.NoSQL
	validator  validator.Validator
	port       Port
	ctxTimeout time.Duration
}

func newGinServer(
	log logger.Logger,
	db repository.NoSQL,
	validator validator.Validator,
	port Port,
	t time.Duration,
) *ginEngine {
	_ = "STUB: not implemented"
	return nil
}

func (g ginEngine) Listen() { _ = "STUB: not implemented"; return }

/* TODO ADD MIDDLEWARE */
func (g ginEngine) setAppHandlers(router *gin.Engine) { _ = "STUB: not implemented"; return }

func (g ginEngine) buildCreateTransferAction() gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

func (g ginEngine) buildFindAllTransferAction() gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

func (g ginEngine) buildCreateAccountAction() gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

func (g ginEngine) buildFindAllAccountAction() gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

func (g ginEngine) buildFindBalanceAccountAction() gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

func (g ginEngine) healthcheck() gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}
