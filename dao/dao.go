package dao

import (
	"context"
	"filter-service/config"
	"golang-kit/db/mysql"
	"golang-kit/log"
)

type Dao struct {
	conf *config.Config
	db   *mysql.DB
}

func NewDao(c *config.Config) (d *Dao, err error) {
	var (
		db *mysql.DB
	)
	d = &Dao{
		conf: c,
	}

	if db, err = mysql.NewMysql(c.Mysql.Master); err != nil {
		log.Error("mysql.NewMysql error(%v)", err)
		return
	} else {
		d.db = db
	}
	return
}

func (d *Dao) Begin(c context.Context) (tx *mysql.Tx, err error) {
	return d.db.Begin(c)
}
