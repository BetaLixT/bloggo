package db

import (
	"testing"

	"github.com/betalixt/bloggo/intl/trace"
	"github.com/betalixt/bloggo/optn"
	"github.com/betalixt/bloggo/util/config"
	"github.com/betalixt/bloggo/util/logger"
	"github.com/spf13/viper"
)

func TestMigration(t *testing.T){
	defer func () {
		if err := recover(); err != nil {
			t.Errorf("failed with: %v", err)
			t.FailNow()
		}
	}()
	lgr := logger.NewLogger()
	berr := config.InitializeConfig(lgr, "../../cfg")
	if berr != nil {
		t.Errorf("failed to create config")
		t.FailNow()
	}
	cfg := viper.GetViper()
	db := NewDatabase(optn.NewDatabaseOptions(cfg))
	tracer := trace.NewZapTracer(lgr)
	trdb := NewTracedDBContext(db, tracer, "test")
	err := RunMigrations(lgr, trdb, nil)
	if err != nil {
		t.Errorf("failed with: %v", err)
		t.FailNow()
	}
}
