package dao

import (
	xsql "database/sql"
	"filter-service/config"
	"github.com/donnie4w/go-logger/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Dao struct {
	conf *config.Config
	db   *sqlx.DB
}

func NewDao(c *config.Config) (d *Dao, err error) {
	d = &Dao{
		conf: c,
	}
	db, err := connectDB(c.Database.Master.Addr)
	d.db = db

	return

}

func connectDB(addr string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open("mysql", addr)
	if err != nil {
		logger.Error(err)
	}
	if err := db.Ping(); err != nil {
		logger.Error(err)
	}
	return
}

func (d *Dao) Begin() (txi *xsql.Tx, err error) {

	txi, err = d.db.Begin()
	if err != nil {
		logger.Error(err)
		return
	}
	return
}
