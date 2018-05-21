package dao

import (
	"context"
	"database/sql"
	"filter-service/model"
	"golang-kit/db/mysql"
	"golang-kit/log"
)

const (
	_addkeywordSQL = "INSERT INTO keyword (content)VALUES(?)"
	_onekeywordSQL = "SELECT id,content FROM keyword WHERE content=?"
)

func (d *Dao) Addkeyword(c context.Context, tx *mysql.Tx, b *model.AddkeywordFrom) (insertId int64, err error) {
	result, err := tx.Exec(_addkeywordSQL, b.Content)
	if err != nil {
		log.Error("d.db.Exec(%+v) error(%v)", b, err)
	}
	return result.LastInsertId()
}

func (d *Dao) GetBykeyword(c context.Context, b *model.AddkeywordFrom) (t *model.KeywordContent, err error) {
	row := d.db.QueryRow(c, _onekeywordSQL, b.Content)
	t = &model.KeywordContent{}

	if err = row.Scan(&t.Id, &t.Content); err != nil {
		if err == sql.ErrNoRows {
			err = nil
			t = nil
			return
		} else {
			log.Error("rows.Scan(%d) rows.Scan error(%v)", &t.Content, err)
			return
		}
	}
	return
}
