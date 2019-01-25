package dao

import (
	"database/sql"
	"fmt"

	"filter-service/model"
	"github.com/donnie4w/go-logger/logger"
)

const (
	_addrelationSQL     = "INSERT INTO relation (keyword_id,state,level,flag)VALUES(?,?,?,?)"
	_uprelationSQL      = "UPDATE  relation SET state =? where id = ?"
	_uprelationlevelSQL = "UPDATE  relation SET level =? where id = ?"
	_delrelationSQL     = "DELETE FROM  relation where id = ?"
	_onerelationSQL     = "SELECT id,keyword_id,state,level,flag FROM relation WHERE keyword_id=? AND flag=?"

	_allrelationListSQL      = "SELECT r.id,r.state,r.level,r.flag,k.content FROM relation as r LEFT JOIN keyword as k ON r.keyword_id = k.id WHERE r.flag = %s AND r.state= %d and level = %d AND k.content like %s order by r.mtime desc limit %d,%d"
	_allrelationListCountSQL = "SELECT count(r.id) as count  FROM relation as r LEFT JOIN keyword as k ON r.keyword_id = k.id WHERE r.flag = %s AND r.state= %d AND level = %d AND k.content like %s"

	_initrelationListSQL = "SELECT r.id,r.state,r.level,r.flag,k.content FROM relation as r LEFT JOIN keyword as k ON r.keyword_id = k.id WHERE r.flag = %s"
)

func (d *Dao) Addrelation(keyword_id, state, level int64, flags string, tx *sql.Tx) (insertId int64, err error) {

	model, err := tx.Prepare(_addrelationSQL)
	if err != nil {
		return
	}
	result, err := model.Exec(keyword_id, state, level, flags)
	insertId, err = result.LastInsertId()

	if err != nil {
		return
	}
	return

}

func (d *Dao) Updaterelation(id, state int64, tx *sql.Tx) (insertId int64, err error) {

	model, err := tx.Prepare(_uprelationSQL)
	if err != nil {
		return
	}

	result, err := model.Exec(state, id)
	insertId, err = result.LastInsertId()

	if err != nil {
		return
	}
	return
}

func (d *Dao) UpdateStatekeyword(id, state int64) (insertId int64, err error) {

	model, err := d.db.Prepare(_uprelationSQL)
	if err != nil {
		return
	}

	result, err := model.Exec(state, id)
	insertId, err = result.LastInsertId()

	if err != nil {
		return
	}
	return
}

func (d *Dao) UpdateLevelkeyword(id, level int64) (insertId int64, err error) {

	model, err := d.db.Prepare(_uprelationlevelSQL)
	if err != nil {
		return
	}

	result, err := model.Exec(level, id)
	insertId, err = result.LastInsertId()

	if err != nil {
		return
	}
	return
}

func (d *Dao) GetByrelation(keyword_id int64, flag string) (r *model.Relation, err error) {

	r = &model.Relation{}
	row := d.db.QueryRow(_onerelationSQL, keyword_id, flag)
	if err = row.Scan(
		&r.Id,
		&r.KeywordId,
		&r.State,
		&r.Level,
		&r.Flag,
	); err != nil {
		if err == sql.ErrNoRows {
			r = nil
			err = nil
		}
		logger.Info(err)
	}
	return
}

func (d *Dao) DelRelationkeyword(id int64) (insertId int64, err error) {

	model, err := d.db.Prepare(_delrelationSQL)
	if err != nil {
		return
	}

	result, err := model.Exec(id)
	insertId, err = result.LastInsertId()

	if err != nil {
		return
	}
	return
}

func (d *Dao) GetByrelationCount(content, flag string, state, level int64) (count int64, err error) {

	row := d.db.QueryRow(fmt.Sprintf(_allrelationListCountSQL, fmt.Sprintf("\"%s\"", flag), state, level, "'%"+content+"%'"))
	if err = row.Scan(&count); err != nil {
		return
	}
	return
}

// //获取列表
func (d *Dao) GetByrelationList(content, flag string, state, level, page, pagesize int64) (rs []*model.Relation, err error) {

	rows, err := d.db.Query(fmt.Sprintf(_allrelationListSQL, fmt.Sprintf("\"%s\"", flag), state, level, "'%"+content+"%'", page, pagesize))

	for rows.Next() {
		r := &model.Relation{}
		if err = rows.Scan(&r.Id, &r.State, &r.Level, &r.Flag, &r.Content); err != nil {
			return
		}
		rs = append(rs, r)
	}
	return
}

// //获取所有关键词
func (d *Dao) GetByrelationKeyword(flag string) (rs []*model.Relation, err error) {

	rows, err := d.db.Query(fmt.Sprintf(_initrelationListSQL, fmt.Sprintf("\"%s\"", flag)))
	for rows.Next() {
		r := &model.Relation{}
		if err = rows.Scan(&r.Id, &r.State, &r.Level, &r.Flag, &r.Content); err != nil {
			return
		}
		rs = append(rs, r)
	}
	return
}
