package dao

import (
	"filter-service/model"
)

const (
	_allListSQL = "SELECT name,id,flag,state FROM business"
)

func (d *Dao) GetByBusinessList() (rs []*model.Business, err error) {
	rows, err := d.db.Query(_allListSQL)
	for rows.Next() {
		r := &model.Business{}
		if err = rows.Scan(&r.Name, &r.Id, &r.Flag, &r.State); err != nil {
			return
		}
		rs = append(rs, r)
	}
	return
}
