package service

import (
	// "fmt"
	// 	"context"
	// 	"filter-service/model"
	// 	"golang-kit/db/mysql"
	"database/sql"
	"filter-service/model"
)

func (s *service) Addkeyword(content string, flag []string, state, level int64) (insertId int64, err error) {

	var (
		tx           *sql.Tx
		keyword_id   int64
		list         *model.KeywordContent
		relationItem *model.Relation
	)
	if tx, err = s.dao.Begin(); err != nil {
		return
	}

	if list, err = s.dao.GetBykeyword(content); err != nil {
		return
	}
	if list == nil {
		if keyword_id, err = s.dao.Addkeyword(content, tx); err != nil {
			tx.Rollback()
			return
		}
	} else {
		keyword_id = list.Id
	}

	for _, flags := range flag {
		if relationItem, err = s.dao.GetByrelation(keyword_id, flags); err != nil {
			return
		}

		if relationItem == nil {
			if _, err = s.dao.Addrelation(keyword_id, state, level, flags, tx); err != nil {
				tx.Rollback()
				return
			}
		} else {
			if _, err = s.dao.Updaterelation(relationItem.Id, state, tx); err != nil {
				tx.Rollback()
				return
			}
		}

	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return
	}
	return
}

func (s *service) adminKeywordList(content, flag string, state, level, page, pagesize int64) (rs []*model.Relation, count int64, err error) {

	var (
		_emplyRs = []*model.Relation{}
	)

	if count, err = s.dao.GetByrelationCount(content, flag, state, level); err != nil {
		return
	}

	pn := (page - 1) * pagesize
	if rs, err = s.dao.GetByrelationList(content, flag, state, level, pn, pagesize); err != nil {
		return
	}

	if len(rs) == 0 {
		rs = _emplyRs
	}
	return
}

func (s *service) UpdateStatekeyword(id, state int64) (insertId int64, err error) {

	if insertId, err = s.dao.UpdateStatekeyword(id, state); err != nil {
		return
	}

	return
}
func (s *service) UpdateLevelkeyword(id, state int64) (insertId int64, err error) {

	if insertId, err = s.dao.UpdateLevelkeyword(id, state); err != nil {
		return
	}

	return
}

func (s *service) DelRelationkeyword(id int64) (insertId int64, err error) {

	if insertId, err = s.dao.DelRelationkeyword(id); err != nil {
		return
	}

	return
}
