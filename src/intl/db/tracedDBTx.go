package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type TracedDBTransaction struct {
  *sqlx.Tx
  lgr *zap.Logger
}

func (tx *TracedDBTransaction) Get (
  dest interface{},
  query string,
  args ...interface{},
) error {
  tx.lgr.Info("Executing query on database")
  err := tx.Tx.Get(dest, query, args...)
  if err != nil {
    tx.lgr.Error("Database query failed", zap.Error(err))
  } else {
    tx.lgr.Info("Database query succeded")
  }
  return err
}

func (tx *TracedDBTransaction) Select (
  dest interface{},
  query string,
  args ...interface{},
) error {
  tx.lgr.Info("Executing query on database")
  err := tx.Tx.Select(dest, query, args...)
  if err != nil {
    tx.lgr.Error("Database query failed", zap.Error(err))
  } else {
    tx.lgr.Info("Database query succeded")
  }
  return err
}

func (tx *TracedDBTransaction) Exec(
	query string,
	args ...interface{},
) (sql.Result, error) {
	tx.lgr.Info("Executing query on database")
	res, err := tx.Tx.Exec(query, args...)	
	if err != nil {
    tx.lgr.Error("Database query failed", zap.Error(err))
  } else {
    tx.lgr.Info("Database query succeded")
  }
  return res, err
}
