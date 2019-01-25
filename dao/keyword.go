package dao

import (
	"database/sql"
	"filter-service/model"
	"github.com/donnie4w/go-logger/logger"
)

const (
	_addkeywordSQL = "INSERT INTO keyword (content)VALUES(?)"
	_onekeywordSQL = "SELECT id,content FROM keyword WHERE content=?"
)

func (d *Dao) Addkeyword(content string, tx *sql.Tx) (insertId int64, err error) {

	model, err := tx.Prepare(_addkeywordSQL)
	if err != nil {
		return
	}
	result, err := model.Exec(content)
	insertId, err = result.LastInsertId()

	if err != nil {
		return
	}
	return

}

func (d *Dao) GetBykeyword(content string) (r *model.KeywordContent, err error) {

	r = &model.KeywordContent{}
	row := d.db.QueryRow(_onekeywordSQL, content)
	if err = row.Scan(
		&r.Id,
		&r.Content,
	); err != nil {
		if err == sql.ErrNoRows {
			r = nil
			err = nil
		}
		logger.Info(err)
	}
	return
}
