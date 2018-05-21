package service

import (
	"context"
	"filter-service/model"
	"golang-kit/db/mysql"
)

func (s *service) Addkeyword(c context.Context, b *model.AddkeywordFrom, flag []string) (insertId int64, err error) {
	var (
		tx           *mysql.Tx
		keyword_id   int64
		keyword      *model.KeywordContent
		relationItem *model.AddkeywordFrom
	)

	if tx, err = s.dao.Begin(c); err != nil {
		return
	}

	if keyword, err = s.dao.GetBykeyword(c, b); err != nil {
		return
	}

	if keyword == nil {
		if keyword_id, err = s.dao.Addkeyword(c, tx, b); err != nil {
			tx.Rollback()
			return
		}
	} else {
		keyword_id = keyword.Id
	}

	for _, flags := range flag {
		relation := &model.AddkeywordFrom{
			Flag:       flags,
			Keyword_id: keyword_id,
		}

		if relationItem, err = s.dao.GetByrelation(c, relation); err != nil {
			return
		}

		if relationItem == nil {
			if _, err = s.dao.Addrelation(c, tx, relation); err != nil {
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

func (s *service) AdminList(c context.Context, content, flag string, pn, ps int64) (bs []*model.Keyword, count int64, err error) {

	var (
		_emplyRs = []*model.Keyword{}
	)

	if count, err = s.dao.GetByrelationCount(c, flag, content); err != nil {
		return
	}

	page := (pn - 1) * ps
	if bs, err = s.dao.GetByrelationList(c, flag, content, page, ps); err != nil {
		return
	}

	if len(bs) == 0 {
		bs = _emplyRs
	}
	return
}

func (s *service) BusinessList(c context.Context) (bs []*model.Business, err error) {
	var (
		_emplyRs = []*model.Business{}
	)

	if bs, err = s.dao.GetByBusinessList(c); err != nil {
		return
	}
	if len(bs) == 0 {
		bs = _emplyRs
	}
	return
}

func (s *service) UpdateStatekeyword(c context.Context, state, id int64) (insertId int64, err error) {
	if insertId, err = s.dao.UpdateStatekeyword(c, state, id); err != nil {
		return
	}
	return
}

func (s *service) UpdateInfokeyword(c context.Context, level, id int64) (insertId int64, err error) {
	if insertId, err = s.dao.UpdateInfokeyword(c, level, id); err != nil {
		return
	}
	return
}
