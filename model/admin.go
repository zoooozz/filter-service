package model

import (
	"github.com/mholt/binding"
	"golang-kit/ecode"
	"golang-kit/log"
	"net/http"
)

type AddkeywordFrom struct {
	Id         int64  `json:"id"`
	Content    string `json:"content"`
	Flag       string `json:"flag"`
	State      int64  `json:"state"`
	Level      int64  `json:"level"`
	Keyword_id int64  `json:"keyword_id"`
}

func (f *AddkeywordFrom) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.Content: "content",
		&f.State:   "state",
		&f.Flag:    "flag",
		&f.Level:   "level",
	}
}

func (f *AddkeywordFrom) Validate(req *http.Request) (err error) {

	if len([]rune(f.Content)) > 100 || len(f.Content) < 1 {
		err = ecode.RequestErr
		log.Error("Pn(%v) is not validate", f.Content)
		return
	}
	return
}
