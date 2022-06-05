package config

import (
	"github.com/betalixt/bloggo/util/blerr"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)


func InitializeConfig (logger *zap.Logger, pth string) (*viper.Viper, *blerr.Error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(pth)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Warn("No config file found")
		} else {
			logger.Error("Failed loading config")
			return nil, blerr.NewError(blerr.ConfigLoadFailureCode, 500, err.Error())
		}
	}

	// TODO bind options
	
	return viper.GetViper(), nil
}

func GetConfig () *viper.Viper {
	return viper.GetViper()
}
