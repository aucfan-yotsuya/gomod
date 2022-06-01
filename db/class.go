package db

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func (d *DB) NewTarget() *Target {
	var tg = new(Target)
	tg.Insert = tg.Do
	d.Target = append(d.Target, tg)
	return tg
}
func (d *DB) NilRedis(index int) bool { return d.Target[index] == nil }
func (d *DB) TargetLen() int          { return len(d.Target) }
func (d *DB) GetTarget(index int) *Target {
	if d.TargetLen() < 1 {
		return nil
	}
	return d.Target[index]
}
func (d *DB) Close() {
	for i := 0; i < d.TargetLen(); i++ {
		d.GetTarget(i).Close()
		d.Target[i] = nil
	}
	d.Target = []*Target{}
}
func (tg *Target) NilConn() bool { return tg.Conn == nil }
func (tg *Target) NilTx() bool   { return tg.Tx == nil }
func (tg *Target) Close() {
	if !tg.NilTx() {
		tg.Tx.Rollback()
		tg.Tx = nil
	}
	if !tg.NilConn() {
		tg.Conn.Close()
		tg.Conn = nil
	}
}
func (tg *Target) NewConn(opt *DbConnOpt) error {
	if tg.Conn, err = sql.Open(opt.Driver, opt.Dsn); err != nil {
		return &Err{Message: err.Error()}
	}
	return nil
}
func (tg *Target) NewPool(opt *DbConnOpt) error {
	tg.NewConn(opt)
	tg.Conn.SetMaxIdleConns(opt.PoolMaxIdle)
	tg.Conn.SetMaxOpenConns(opt.PoolMaxActive)
	tg.Conn.SetConnMaxIdleTime(opt.PoolConnMaxIdleTime)
	tg.Conn.SetConnMaxLifetime(opt.PoolConnMaxLifetime)
	return nil
}
func (tg *Target) Begin() error {
	if tg.Tx, err = tg.Conn.Begin(); err != nil {
		return &Err{Message: err.Error()}
	}
	return nil
}
func (tg *Target) Commit() error {
	if tg.NilTx() {
		return &Err{Message: "Tx has nil"}
	}
	if err = tg.Tx.Commit(); err != nil {
		return &Err{Message: err.Error()}
	}
	return nil
}
func (tg *Target) Rollback() error {
	if tg.NilTx() {
		return &Err{Message: "Tx has nil"}
	}
	if err = tg.Tx.Rollback(); err != nil {
		return &Err{Message: err.Error()}
	}
	return nil
}
func (tg *Target) Do(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	var r sql.Result
	if r, err = tg.Conn.ExecContext(ctx, query, args...); err != nil {
		return r, &Err{Message: err.Error()}
	}
	return r, nil
}
func (tg *Target) DoTx(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	var r sql.Result
	if r, err = tg.Tx.ExecContext(ctx, query, args...); err != nil {
		return r, &Err{Message: err.Error()}
	}
	return r, nil
}
func (tg *Target) Select(ctx context.Context, query string, args ...interface{}) ([]map[string][]byte, error) {
	var r *sql.Rows
	if r, err = tg.Conn.QueryContext(ctx, query, args...); err != nil {
		if IsNoRows(err) {
			return []map[string][]byte{}, err
		} else {
			return []map[string][]byte{}, &Err{Message: err.Error()}
		}
	}
	defer r.Close()

	var (
		values   = make([]sql.RawBytes, func() int { c, _ := r.Columns(); return len(c) }())
		scanArgs = make([]interface{}, len(values))
		rows     = []map[string][]byte{}
		columns  []string
	)
	if columns, err = r.Columns(); err != nil {
		return rows, &Err{Message: err.Error()}
	}

	for i := range values {
		scanArgs[i] = &values[i]
	}
	for r.Next() {
		if err = r.Scan(scanArgs...); err != nil {
			return rows, &Err{Message: err.Error()}
		}
		var row = make(map[string][]byte)
		for i, col := range values {
			row[columns[i]] = make([]byte, len(col))
			copy(row[columns[i]], col)
		}
		rows = append(rows, row)
	}
	return rows, nil
}
func (tg *Target) BulkInsert(ctx context.Context, query string, args []interface{}) (sql.Result, error) {
	var r sql.Result
	if r, err = tg.Conn.ExecContext(ctx, query, args...); err != nil {
		return r, &Err{Message: err.Error()}
	}
	return r, nil
}
func (tg *Target) BulkInsertTx(ctx context.Context, query string, args []interface{}) (sql.Result, error) {
	var r sql.Result
	if r, err = tg.Tx.ExecContext(ctx, query, args...); err != nil {
		return r, &Err{Message: err.Error()}
	}
	return r, nil
}
