package service

import (
	"filter-service/model"
	"github.com/mholt/binding"
	"golang-kit/ecode"
	"golang-kit/log"
	"golang-kit/net/context"
	"strconv"
	"strings"
)

func internalAddkeyword(c context.Context) {
	var (
		err  error
		flag []string
	)
	result := c.Result()
	form := &model.AddkeywordFrom{}
	if err = binding.Bind(c.Request(), form); err != nil {
		log.Error("binding.Bind error(%v)", err)
		err = ecode.RequestErr
		result["code"] = err
		return
	}

	params := c.Request().Form
	flagStr := params.Get("flag")
	if flagStr == "" {
		flag = []string{"all"}
	} else {
		flag = strings.Split(flagStr, ",")
	}

	levelStr := params.Get("level")
	if levelStr != "" {
		if form.Level, err = strconv.ParseInt(levelStr, 10, 64); err != nil {
			log.Error("strconv.ParseInt(%s) error(%v)", levelStr, err)
			result["code"] = ecode.RequestErr
			return
		}
	}
	stateStr := params.Get("state")
	if stateStr != "" {
		if form.Level, err = strconv.ParseInt(stateStr, 10, 64); err != nil {
			log.Error("strconv.ParseInt(%s) error(%v)", stateStr, err)
			result["code"] = ecode.RequestErr
			return
		}
	}

	if _, err = svr.Addkeyword(c, form, flag); err != nil {
		result["code"] = err
		return
	}
	result["code"] = ecode.OK
	return

}

func internalList(c context.Context) {
	var (
		err     error
		flag    string
		content string
		pn      int64
		ps      int64
		bs      []*model.Keyword
		count   int64
	)
	result := c.Result()
	params := c.Request().Form
	flag = params.Get("flag")
	if flag == "" {
		result["code"] = ecode.RequestErr
		return
	}
	content = params.Get("content")

	psStr := params.Get("ps")
	pnStr := params.Get("pn")

	pn, err = strconv.ParseInt(pnStr, 10, 64)
	if err != nil || pn < 1 {
		pn = 1
	}
	ps, err = strconv.ParseInt(psStr, 10, 64)
	if err != nil || ps < 0 || ps > 200 {
		ps = 100
	}

	if bs, count, err = svr.AdminList(c, content, flag, pn, ps); err != nil {
		result["code"] = err
		return
	}
	result["data"] = map[string]interface{}{
		"bs":    bs,
		"count": count,
	}
	result["code"] = ecode.OK
	return
}

func internalBusinessList(c context.Context) {
	var (
		err error
		bs  []*model.Business
	)
	result := c.Result()
	if bs, err = svr.BusinessList(c); err != nil {
		result["code"] = err
		return
	}
	result["data"] = map[string]interface{}{
		"bs": bs,
	}

	result["code"] = ecode.OK
	return
}

func internalUpdateStatekeyword(c context.Context) {
	var (
		err   error
		state int64
		id    int64
	)
	result := c.Result()
	params := c.Request().Form
	stateStr := params.Get("state")
	idStr := params.Get("id")

	if state, err = strconv.ParseInt(stateStr, 10, 64); err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", stateStr, err)
		result["code"] = ecode.RequestErr
		return
	}

	if id, err = strconv.ParseInt(idStr, 10, 64); err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", idStr, err)
		result["code"] = ecode.RequestErr
		return
	}

	if state != 0 && state != 1 {
		result["code"] = ecode.RequestErr
		return
	}
	if _, err = svr.UpdateStatekeyword(c, state, id); err != nil {
		result["code"] = err
		return
	}
	result["code"] = ecode.OK
	return
}

func internalUpdateInfokeyword(c context.Context) {
	var (
		err   error
		level int64
		id    int64
	)
	result := c.Result()
	params := c.Request().Form
	levelStr := params.Get("level")
	idStr := params.Get("id")

	if level, err = strconv.ParseInt(levelStr, 10, 64); err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", levelStr, err)
		result["code"] = ecode.RequestErr
		return
	}

	if id, err = strconv.ParseInt(idStr, 10, 64); err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", idStr, err)
		result["code"] = ecode.RequestErr
		return
	}

	if _, err = svr.UpdateInfokeyword(c, level, id); err != nil {
		result["code"] = err
		return
	}
	result["code"] = ecode.OK
	return
}
