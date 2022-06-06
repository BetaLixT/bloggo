package main

import (
	"github.com/betalixt/bloggo/mw"
	"github.com/betalixt/bloggo/optn"
	"github.com/betalixt/bloggo/svc"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func NewGinEngine(
	lgr *zap.Logger,
	corsOptn *optn.CorsOptions,
	tknSvc *svc.TokenService,
	db *sqlx.DB,
) *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.SetTrustedProxies(nil)

	// - Setting up middlewares
	router.Use(mw.TransactionContextGenerationMiddleware(lgr, db))
	router.Use(mw.LoggingMiddleware())
	router.Use(mw.RecoveryMiddleware(lgr))
	router.Use(mw.CorsMiddleware(lgr, corsOptn))
  // TODO Make this configurable
	router.Use(mw.AuthMiddleware(tknSvc))

	// - Responding to head
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "alive",
		})
	})

	return router
}
