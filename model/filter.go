package model

import (
	"filter-service/actire"
	"time"
)

type KeywordContent struct {
	Id      int64     `json:"id"`
	Content string    `json:"content"`
	Ctime   time.Time `json:"ctime"`
	Mtime   time.Time `json:"mtime"`
}

type Relation struct {
	Id        int64     `json:"id"`
	Content   string    `json:"content"`
	Flag      string    `json:"flag"`
	State     int64     `json:"state"`
	Level     int64     `json:"level"`
	KeywordId int64     `json:"keyword_id"`
	Ctime     time.Time `json:"ctime"`
	Mtime     time.Time `json:"mtime"`
}

type Filter struct {
	Matcher *actire.Matcher
}

type Business struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Flag  string `json:"flag"`
	State string `json:"state"`
}
