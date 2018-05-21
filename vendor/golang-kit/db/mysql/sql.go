package mysql

import (
	"context"
	xsql "database/sql"
	"errors"
	"fmt"
	opentracing "github.com/opentracing/opentracing-go"
)

var (
	ErrStmtNil = errors.New("prepare failed and stmt nil")
	ErrNoRows  = xsql.ErrNoRows
	ErrTxDone  = xsql.ErrTxDone
)

type DB struct {
	*xsql.DB
	addr string
}

type Tx struct {
	db   *DB
	tx   *xsql.Tx
	addr string
	span opentracing.Span
}

type Row struct {
	row  *xsql.Row
	span opentracing.Span
}

type Rows struct {
	*xsql.Rows
}

func Open(driverName, dataSourceName string) (*DB, error) {
	d, err := xsql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{DB: d}, nil
}

func (db *DB) Begin(c context.Context) (tx *Tx, err error) {
	var (
		span opentracing.Span
		txi  *xsql.Tx
	)
	if txi, err = db.DB.Begin(); err != nil {
		return
	}
	if c.Value("sync") == true {
		span, _ = opentracing.StartSpanFromContext(c, "mysql begin")
	} else {
		span = nil
	}

	return &Tx{tx: txi, span: span, addr: db.addr, db: db}, nil
}

func (db *DB) Exec(c context.Context, query string, args ...interface{}) (res xsql.Result, err error) {
	if c.Value("sync") == true {
		span, _ := opentracing.StartSpanFromContext(c, fmt.Sprintf("mysql %s", query))
		span.LogEvent(fmt.Sprintf("mysql %s %v", query, args))
		defer span.Finish()
	}
	res, err = db.DB.Exec(query, args...)
	return
}

func (db *DB) Ping(c context.Context) (err error) {
	if c.Value("sync") == true {
		span, _ := opentracing.StartSpanFromContext(c, "mysql ping")
		defer span.Finish()
	}
	err = db.DB.Ping()
	return
}

func (db *DB) Query(c context.Context, query string, args ...interface{}) (rows *Rows, err error) {
	var rrows *xsql.Rows
	if c.Value("sync") == true {
		span, _ := opentracing.StartSpanFromContext(c, fmt.Sprintf("mysql %s", query))
		span.LogEvent(fmt.Sprintf("mysql %s", query))
		defer span.Finish()
	}
	if rrows, err = db.DB.Query(query, args...); err == nil {
		rows = &Rows{rrows}
	}
	return
}

func (db *DB) QueryRow(c context.Context, query string, args ...interface{}) *Row {
	var span opentracing.Span
	if c.Value("sync") == true {
		span, _ = opentracing.StartSpanFromContext(c, fmt.Sprintf("mysql %s", query))
		span.LogEvent(fmt.Sprintf("mysql %s %v", query, args))
	} else {
		span = nil
	}
	row := db.DB.QueryRow(query, args...)
	return &Row{row: row, span: span}
}

func (tx *Tx) Commit() (err error) {
	if tx.span != nil {
		defer tx.span.Finish()
	}
	err = tx.tx.Commit()
	return
}

func (tx *Tx) Rollback() (err error) {
	if tx.span != nil {
		defer tx.span.Finish()
	}
	err = tx.tx.Rollback()
	return
}

func (tx *Tx) Exec(query string, args ...interface{}) (res xsql.Result, err error) {
	res, err = tx.tx.Exec(query, args...)
	return
}

func (r *Row) Scan(dest ...interface{}) (err error) {
	if r.span != nil {
		defer r.span.Finish()
	}
	if r.row != nil {
		err = r.row.Scan(dest...)
	} else {
		err = ErrStmtNil
	}
	return
}
