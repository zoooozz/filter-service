package model

import (
	"github.com/mholt/binding"
	"golang-kit/ecode"
	"golang-kit/log"
	"net/http"
)

// add
func (f *ClockTask) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.BusinessId: "business_id",
	}
}

func (f *ClockTask) Validate(req *http.Request) (err error) {
	if f.BusinessId == int64(0) {
		err = ecode.RequestErr
		log.Error("BusinessId(%v) is not validate", f.BusinessId)
	}
	return
}

// sel
type SelBusinessTaskForm struct {
	Pn int64
	Ps int64
}

func (f *SelBusinessTaskForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.Pn: "pn",
		&f.Ps: "ps",
	}
}

func (f *SelBusinessTaskForm) Validate(req *http.Request) (err error) {
	if f.Pn < 1 {
		err = ecode.RequestErr
		log.Error("Pn(%v) is not validate", f.Pn)
		return
	}
	if f.Ps < 1 {
		err = ecode.RequestErr
		log.Error("Ps(%v) is not validate", f.Ps)
		return
	}
	return
}
