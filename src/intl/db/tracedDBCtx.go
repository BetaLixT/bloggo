package db

import (
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
  err := trDB.DB.Get(dest, query, args)
  if err != nil {
    trDB.lgr.Error("Database query failed")
  } else {
    trDB.lgr.Info("Database query succeded")
  }
  return err
}
