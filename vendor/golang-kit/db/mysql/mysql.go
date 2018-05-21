package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"golang-kit/config"
	"golang-kit/log"
)

func NewMysql(c *config.Mysql) (db *DB, err error) {
	if db, err = Open("mysql", c.DSN); err != nil {
		log.Error("open mysql error(%v)", err)
		return
	}
	db.addr = c.Addr
	db.SetMaxOpenConns(c.Active)
	db.SetMaxIdleConns(c.Idle)
	return
}
