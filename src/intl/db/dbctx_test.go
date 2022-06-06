package db

import (
	"testing"

	"github.com/betalixt/bloggo/optn"
	"github.com/betalixt/bloggo/util/config"
	"github.com/betalixt/bloggo/util/logger"
	"github.com/spf13/viper"
)

func TestDatabaseConnection(t *testing.T){
	defer func () {
		if err := recover(); err != nil {
			t.Errorf("failed with: %v", err)
			t.FailNow()
		}
	}()
	lgr := logger.NewLogger()
	err := config.InitializeConfig(lgr, "../../cfg")
	if err != nil {
		t.Errorf("failed to create config")
		t.FailNow()
	}
	cfg := viper.GetViper()
	_ = NewDatabase(optn.NewDatabaseOptions(cfg))
}
