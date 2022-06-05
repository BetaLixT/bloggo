package config

import (
	"strings"

	"github.com/betalixt/bloggo/util/blerr"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)


func InitializeConfig (
	logger *zap.Logger,
	pth string,
) *blerr.Error {
	viper.SetConfigName("config")
	viper.AddConfigPath(pth)
	viper.SetEnvPrefix("BLOGGO")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()
	viper.BindEnv("PORT", "PORT")


	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Warn("No config file found")
		} else {
			logger.Error("Failed loading config")
			return blerr.NewError(blerr.ConfigLoadFailureCode, 500, err.Error())
		}
	}
	
	return nil
}

func NewConfig (lgr *zap.Logger) *viper.Viper {
	if err := InitializeConfig(lgr, "./cfg"); err != nil {
		panic(err)
	}

	return viper.GetViper()
}
