package model

import (
	"github.com/mholt/binding"
	"golang-kit/ecode"
	"golang-kit/log"
	"golang-kit/time"
	"net/http"
)

// add
func (f *Template) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.Name:      "name",
		&f.Content:   "content",
		&f.Cover:     "cover",
		&f.Paragraph: "paragraph",
		&f.Sentence:  "sentence",
	}
}

func (f *Template) Validate(req *http.Request) (err error) {
	if f.Name == "" {
		err = ecode.RequestErr
		log.Error("Name(%v) is not validate", f.Name)
		return
	}
	if f.Content == "" {
		err = ecode.RequestErr
		log.Error("Content(%v) is not validate", f.Content)
		return
	}
	if f.Cover == "" {
		err = ecode.RequestErr
		log.Error("Cover(%v) is not validate", f.Cover)
		return
	}
	if f.Paragraph == "" {
		err = ecode.RequestErr
		log.Error("Paragraph(%v) is not validate", f.Paragraph)
		return
	}
	if f.Sentence == "" {
		err = ecode.RequestErr
		log.Error("Sentence(%v) is not validate", f.Sentence)
		return
	}
	return
}

// sel template
type SelTemplateForm struct {
	Pn int64
	Ps int64
}

func (f *SelTemplateForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.Pn: "pn",
		&f.Ps: "ps",
	}
}

func (f *SelTemplateForm) Validate(req *http.Request) (err error) {
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

type SelSendedForm struct {
	Pn int64
	Ps int64
}

func (f *SelSendedForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.Pn: "pn",
		&f.Ps: "ps",
	}
}

func (f *SelSendedForm) Validate(req *http.Request) (err error) {
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

// user
type UserListForm struct {
	Pn      int64
	Ps      int64
	Account string
	Uname   string
}

func (f *UserListForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.Pn:      "pn",
		&f.Ps:      "ps",
		&f.Account: "account",
		&f.Uname:   "uname",
	}
}

func (f *UserListForm) Validate(req *http.Request) (err error) {
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

// 修改用户服务信息
type UserSvrForm struct {
	Uid       int64
	Level     int8
	StartTime time.Time
	EndTime   time.Time
	Start     int64
	End       int64
}

func (f *UserSvrForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.Uid:   "uid",
		&f.Level: "level",
		&f.Start: "start",
		&f.End:   "end",
	}
}

func (f *UserSvrForm) Validate(req *http.Request) (err error) {
	if f.Level < 1 {
		err = ecode.RequestErr
		log.Error("Level(%v) is not validate", f.Level)
		return
	}
	f.StartTime = time.Time(f.Start)
	f.EndTime = time.Time(f.End)
	return
}
