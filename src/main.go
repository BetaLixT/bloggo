package main

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/betalixt/bloggo/svc"
	"github.com/betalixt/bloggo/util/blerr"
	"github.com/betalixt/bloggo/util/config"
	"github.com/betalixt/bloggo/util/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	// - Custom panic logging
	defer func() {
		if err := recover(); err != nil {	
			fmt.Printf("[Panic Encountered] %v\n", time.Now())
			prsd, ok := err.(blerr.Error); if ok {
				fmt.Printf(
					"ErrorCode: %d [%s]\n",
					prsd.ErrorCode,
					prsd.ErrorMessage,
				)
				if prsd.ErrorDetail != "" {
					fmt.Printf("Details: %s\n", prsd.ErrorDetail)
				}
			} else {
				fmt.Printf("error: %v\n", err)
			}
			fmt.Printf("%s", string(debug.Stack()))
		}
	}()

	app := fx.New(
		fx.Provide(logger.NewLogger),
		fx.Provide(config.NewConfig),
		// fx.Provide(db.NewDatabase),
		// fx.Provide(db.NewMigration),
		fx.Provide(svc.NewTokenService),
		fx.Provide(NewGinEngine),
		fx.Invoke(StartService),
	)
	app.Run();
}

func StartService(
	cfg *viper.Viper,
	lgr *zap.Logger,
	// mgr *db.Migration,
	gin *gin.Engine) {
	
	port := cfg.GetString("PORT")
	if port == "" {
		lgr.Warn("No port was specified, using 8080")
		port = "8080"
	} else if _, err := strconv.Atoi(port); err != nil {
		lgr.Error("Non numeric value was specified for port")
		panic(blerr.NewError(blerr.InvalidPortCode, 500, ""))
	}

	// lgr.Info("Running migrations")
	// err := mgr.RunMigrations()
	// if err != nil {
	// 	lgr.Warn("Failed running migrations", zap.Error(err))
	// }
	
	lgr.Info("Starting service", zap.String("port", port))
	gin.Run(fmt.Sprintf(":%s", port))
}
