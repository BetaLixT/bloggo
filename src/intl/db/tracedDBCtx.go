package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type TracedDBContext struct {
  *sqlx.DB
  lgr *zap.Logger
}

func NewTracedDBContext(db *sqlx.DB, lgr *zap.Logger) *TracedDBContext {
  
  return &TracedDBContext{
    lgr: lgr,
    DB: db,
  }
}

func (trDB *TracedDBContext) Get (
  dest interface{},
  query string,
  args ...interface{},
) error {
  trDB.lgr.Info("Executing query on database")
  err := trDB.DB.Get(dest, query, args...)
  if err != nil {
    trDB.lgr.Error("Database query failed", zap.Error(err))
  } else {
    trDB.lgr.Info("Database query succeded")
  }
  return err
}

func (trDB *TracedDBContext) Select (
  dest interface{},
  query string,
  args ...interface{},
) error {
  trDB.lgr.Info("Executing query on database")
  err := trDB.DB.Select(dest, query, args...)
  if err != nil {
    trDB.lgr.Error("Database query failed", zap.Error(err))
  } else {
    trDB.lgr.Info("Database query succeded")
  }
  return err
}

func (db *TracedDBContext) Exec(
	query string,
	args ...interface{},
) (sql.Result, error) {
	db.lgr.Info("Executing query on database")
	res, err := db.DB.Exec(query, args...)	
	if err != nil {
    db.lgr.Error("Database query failed", zap.Error(err))
  } else {
    db.lgr.Info("Database query succeded")
  }
  return res, err
}

func (db *TracedDBContext) Beginx () (*TracedDBTransaction, error) {
  tx, err := db.DB.Beginx()
  return &TracedDBTransaction{
    Tx: tx,
    lgr: db.lgr,
  }, err
}

func (db *TracedDBContext) MustBegin () *TracedDBTransaction {
  tx := db.DB.MustBegin()
  return &TracedDBTransaction{
    Tx: tx,
    lgr: db.lgr,
  }
}
