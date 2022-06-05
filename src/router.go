package main

import (
	"github.com/betalixt/bloggo/mw"
	"github.com/betalixt/bloggo/svc"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewGinEngine(
	lgr *zap.Logger,
	cfg *viper.Viper,
	tknSvc *svc.TokenService,
) *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.SetTrustedProxies(nil)

	// - Setting up middlewares
	router.Use(mw.TransactionContextGenerationMiddleware(lgr))
	router.Use(mw.LoggingMiddleware())
	router.Use(mw.RecoveryMiddleware(lgr))
	router.Use(mw.CorsMiddleware(lgr, cfg))
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
