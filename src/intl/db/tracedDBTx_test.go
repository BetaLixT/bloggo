package db

import (
	"testing"

	"github.com/betalixt/bloggo/intl/trace"
	"github.com/betalixt/bloggo/optn"
	"github.com/betalixt/bloggo/util/config"
	"github.com/betalixt/bloggo/util/logger"
	"github.com/spf13/viper"
)

func TestGet(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("failed with %v", err)
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
  tdb := NewTracedDBContext(db, tracer, "test")
  tx := tdb.MustBegin()
  chck := ExistsEntity{}
  err := tx.Get(&chck, CheckTimestampProceduresExist)
  tx.Commit()
  if err != nil {
    t.Errorf("failed with %v", err)
    t.FailNow()
  }
}
