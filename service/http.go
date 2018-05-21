package service

import (
	"golang-kit/ecode"
	"golang-kit/log"
	"golang-kit/net/context"
)

func list(c context.Context) {
	var (
		err   error
		t     string
		k     []string
		level int64
	)
	result := c.Result()
	params := c.Request().Form

	content := params.Get("content")
	if content == "" {
		log.Error("strconv.ParseInt(%s) error(%v)", content, err)
		result["code"] = ecode.RequestErr
		return
	}
	flag := params.Get("flag")
	if content == "" {
		log.Error("strconv.ParseInt(%s) error(%v)", flag, err)
		result["code"] = ecode.RequestErr
		return

	}

	if t, level, k, err = svr.KeywordFilter(c, content, flag); err != nil {
		result["code"] = err
		return
	}

	result["data"] = map[string]interface{}{
		"content": t,
		"keyword": k,
		"level":   level,
	}
	result["code"] = ecode.OK
	return
}
