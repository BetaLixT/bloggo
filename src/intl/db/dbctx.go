package db

import (	
	"github.com/jmoiron/sqlx"
	"github.com/betalixt/bloggo/util/blerr"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func NewDatabase(cfg *viper.Viper) *sqlx.DB {

	// TODO Move this to options
	conn := cfg.GetString("DatabaseOptions.ConnectionString")
	db, err := sqlx.Open("postgres", conn)
	if err != nil {
		panic(blerr.NewError(
				blerr.DatabaseConnectionOpenFailure,
				500,
				err.Error(),
			))
	}

	err = db.Ping()
	if err != nil {
		panic(blerr.NewError(blerr.DatabasePingFailure, 500, err.Error()))
	}

	return db
}
