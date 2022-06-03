package db

import (
	"context"
	"database/sql"
	"time"
)

type (
	DB struct {
		Target []*Target
	}
	Target struct {
		Conn     *sql.DB
		Tx       *sql.Tx
		Insert   func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
		InsertTx func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	}
	DbConnOpt struct {
		Driver, Dsn                                       string
		PoolMaxActive, PoolMaxIdle                        int
		Timeout, PoolConnMaxIdleTime, PoolConnMaxLifetime time.Duration
	}
)

var (
	err error
	d   *DB
	tg  *Target
)

func New() *DB {
	return new(DB)
}
