package model

import (
	"filter-service/actire"
	"golang-kit/time"
)

type KeywordContent struct {
	Id      int64     `json:"id"`
	Content string    `json:"content"`
	Ctime   time.Time `json:"ctime"`
	Mtime   time.Time `json:"mtime"`
}

type Keyword struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
	Flag    string `json:"flag"`
	State   int64  `json:"state"`
	Level   int64  `json:"level"`
}

type Business struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Flag  string `json:"flag"`
	State string `json:"state"`
}

type Filter struct {
	Matcher *actire.Matcher
}
