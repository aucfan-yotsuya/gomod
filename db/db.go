package db

import (
	"database/sql"
	"time"
)

type (
	DB struct {
		Target []*Target
	}
	Target struct {
		Conn *sql.DB
		Tx   *sql.Tx
	}
	DbConnOpt struct {
		Driver, Dsn                                       string
		Port, PoolMaxActive, PoolMaxIdle                  int
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
