package mw

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func CorsMiddleware(lgr *zap.Logger, cfg *viper.Viper) gin.HandlerFunc {
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowOrigins = cfg.GetStringSlice("AllowedOrigins")
	corsCfg.AllowCredentials = true
	corsCfg.AllowHeaders = []string {
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization",
		"accept",
		"origin",
		"Cache-Control",
		"X-Requested-With",
	}
	lgr.Info(
		"Configuring cors",
		zap.Strings("allowedOrigins", corsCfg.AllowOrigins),
	)
	return cors.New(corsCfg)
}