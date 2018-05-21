package model

import (
	"golang-kit/time"
)

type SysNotice struct {
	Id      int64     `json:"id"`
	Title   string    `json:"title"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Content string    `json:"content"`
	Stime   int64     `json:"stime"`
}
