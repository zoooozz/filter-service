package dao

import (
	"context"
	"database/sql"
	"filter-service/model"
	"fmt"
	"golang-kit/db/mysql"
	"golang-kit/log"
)

const (
	_addrelationSQL          = "INSERT INTO relation (keyword_id,state,level,flag)VALUES(?,?,?,?)"
	_uprelationSQL           = "UPDATE  relation SET state =? where id = ?"
	_uprelationlevelSQL      = "UPDATE  relation SET level =? where id = ?"
	_onerelationSQL          = "SELECT id,keyword_id,state,level,flag FROM relation WHERE keyword_id=? AND flag=?"
	_allrelationListSQL      = "SELECT r.id,r.state,r.level,r.flag,k.content FROM relation as r LEFT JOIN keyword as k ON r.keyword_id = k.id WHERE r.flag = %s AND k.content like %s limit %d,%d"
	_allrelationListCountSQL = "SELECT count(r.id) as count  FROM relation as r LEFT JOIN keyword as k ON r.keyword_id = k.id WHERE r.flag = %s AND k.content like %s"
	_initrelationListSQL     = "SELECT r.id,r.state,r.level,r.flag,k.content FROM relation as r LEFT JOIN keyword as k ON r.keyword_id = k.id WHERE r.flag = %s"
)

//新增关键词关系
func (d *Dao) Addrelation(c context.Context, tx *mysql.Tx, b *model.AddkeywordFrom) (insertId int64, err error) {
	result, err := tx.Exec(_addrelationSQL, b.Keyword_id, b.State, b.Level, b.Flag)
	if err != nil {
		log.Error("d.db.Exec(%+v) error(%v)", b, err)
	}
	return result.RowsAffected()
}

//获取关系
func (d *Dao) GetByrelation(c context.Context, b *model.AddkeywordFrom) (t *model.AddkeywordFrom, err error) {
	row := d.db.QueryRow(c, _onerelationSQL, b.Keyword_id, b.Flag)
	t = &model.AddkeywordFrom{}
	if err = row.Scan(&t.Id, &t.Keyword_id, &t.State, &t.Level, &t.Flag); err != nil {
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

//获取列表数量
func (d *Dao) GetByrelationCount(c context.Context, flag, content string) (count int64, err error) {

	row := d.db.QueryRow(c, fmt.Sprintf(_allrelationListCountSQL, fmt.Sprintf("\"%s\"", flag), "'%"+content+"%'"))
	if err = row.Scan(&count); err != nil {
		log.Error("row.Scan() err(%v)", err)
	}
	return
}

//获取列表
func (d *Dao) GetByrelationList(c context.Context, flag, content string, pn, ps int64) (ts []*model.Keyword, err error) {

	rows, err := d.db.Query(c, fmt.Sprintf(_allrelationListSQL, fmt.Sprintf("\"%s\"", flag), "'%"+content+"%'", pn, ps))
	for rows.Next() {
		t := &model.Keyword{}
		if err = rows.Scan(&t.Id, &t.State, &t.Level, &t.Flag, &t.Content); err != nil {
			log.Error("rows.Scan(%d) rows.Scan error(%v)", t, err)
			return
		}
		ts = append(ts, t)
	}
	return
}

//更新关系状态
func (d *Dao) UpdateStatekeyword(c context.Context, state, id int64) (insertId int64, err error) {
	result, err := d.db.Exec(c, _uprelationSQL, state, id)
	if err != nil {
		log.Error("d.db.Exec(%+v) error(%v)", id, err)
	}
	return result.RowsAffected()
}

//更新关系信息
func (d *Dao) UpdateInfokeyword(c context.Context, level, id int64) (insertId int64, err error) {
	result, err := d.db.Exec(c, _uprelationlevelSQL, level, id)
	if err != nil {
		log.Error("d.db.Exec(%+v) error(%v)", id, err)
	}
	return result.RowsAffected()
}

//获取所有关键词
func (d *Dao) GetByrelationKeyword(c context.Context, flag string) (ts []*model.Keyword, err error) {

	rows, err := d.db.Query(c, fmt.Sprintf(_initrelationListSQL, fmt.Sprintf("\"%s\"", flag)))
	for rows.Next() {
		t := &model.Keyword{}
		if err = rows.Scan(&t.Id, &t.State, &t.Level, &t.Flag, &t.Content); err != nil {
			log.Error("rows.Scan(%d) rows.Scan error(%v)", t, err)
			return
		}
		ts = append(ts, t)
	}
	return
}
