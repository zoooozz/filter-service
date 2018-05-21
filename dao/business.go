package dao

import (
	"context"
	"filter-service/model"
	"golang-kit/log"
)

const (
	_allListSQL = "SELECT name,id,flag,state FROM business"
)

func (d *Dao) GetByBusinessList(c context.Context) (ts []*model.Business, err error) {
	rows, err := d.db.Query(c, _allListSQL)
	for rows.Next() {
		t := &model.Business{}
		if err = rows.Scan(&t.Name, &t.Id, &t.Flag, &t.State); err != nil {
			log.Error("rows.Scan(%d) rows.Scan error(%v)", t, err)
			return
		}
		ts = append(ts, t)
	}
	return
}
