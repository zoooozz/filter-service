package service

import (
	"filter-service/model"
	"github.com/labstack/echo"
	"strconv"
	"strings"
)

func addKeyword(c echo.Context) error {

	var (
		flag    []string
		level   int64
		state   int64
		err     error
		content string
	)
	flagStr := c.FormValue("flag")
	if flagStr == "" {
		flag = []string{"all"}
	} else {
		flag = strings.Split(flagStr, ",")
	}

	levelStr := c.FormValue("level")
	if levelStr != "" {
		if level, err = strconv.ParseInt(levelStr, 10, 64); err != nil {
			return c.JSON(model.OutputRet(model.Notdefinitionparams))
		}

		if level != model.LevelLight && level != model.LevelSevere {
			return c.JSON(model.OutputRet(model.Notdefinitionparams))
		}
	} else {
		level = 1
	}

	stateStr := c.FormValue("state")
	if stateStr != "" {

		if state, err = strconv.ParseInt(stateStr, 10, 64); err != nil {
			return c.JSON(model.OutputRet(model.Notdefinitionparams))
		}

		if state != model.StateOpen && state != model.StateClose {
			return c.JSON(model.OutputRet(model.Notdefinitionparams))
		}

	} else {
		state = 1
	}

	content = c.FormValue("content")
	if len([]rune(content)) > 100 || len(content) < 1 {
		return c.JSON(model.OutputRet(model.Notdefinitionparams))
	}

	_, err = svr.Addkeyword(content, flag, state, level)
	if err != nil {
		return c.JSON(model.OutputRet(model.RetErr))
	}

	result := map[string]interface{}{
		"list": "",
	}
	return c.JSON(model.OutputRet(model.Success, result))
}

func editKeyword(c echo.Context) error {

	var (
		err   error
		state int64
		level int64
		id    int64
	)
	stateStr := c.FormValue("state")
	levelStr := c.FormValue("level")
	idStr := c.FormValue("id")

	if stateStr != "" {
		if state, err = strconv.ParseInt(stateStr, 10, 64); err != nil {
			return c.JSON(model.OutputRet(model.Notdefinitionparams))
		}

		if state != model.StateClose && state != model.StateOpen {
			return c.JSON(model.OutputRet(model.Notdefinitionparams))
		}

	}

	if levelStr != "" {
		if level, err = strconv.ParseInt(levelStr, 10, 64); err != nil {
			return c.JSON(model.OutputRet(model.Notdefinitionparams))
		}

		if level != model.LevelLight && level != model.LevelSevere {
			return c.JSON(model.OutputRet(model.Notdefinitionparams))
		}
	}

	if stateStr == "" && levelStr == "" {
		return c.JSON(model.OutputRet(model.Notdefinitionparams))
	}

	if id, err = strconv.ParseInt(idStr, 10, 64); err != nil {
		return c.JSON(model.OutputRet(model.Notdefinitionparams))
	}

	if stateStr != "" {
		_, err = svr.UpdateStatekeyword(id, state)
	}

	if levelStr != "" {
		_, err = svr.UpdateLevelkeyword(id, level)

	}
	if err != nil {
		return c.JSON(model.OutputRet(model.RetErr))
	}

	result := map[string]interface{}{
		"list": "",
	}
	return c.JSON(model.OutputRet(model.Success, result))

}

func delKeyword(c echo.Context) error {
	var (
		err error
		id  int64
	)

	idStr := c.FormValue("id")
	if id, err = strconv.ParseInt(idStr, 10, 64); err != nil {
		return c.JSON(model.OutputRet(model.Notdefinitionparams))
	}

	_, err = svr.DelRelationkeyword(id)

	if err != nil {
		return c.JSON(model.OutputRet(model.RetErr))
	}

	result := map[string]interface{}{
		"list": "",
	}
	return c.JSON(model.OutputRet(model.Success, result))
}

func adminKeywordList(c echo.Context) error {

	var (
		err      error
		flag     string
		content  string
		level    int64
		state    int64
		page     int64
		pagesize int64
		ret      []*model.Relation
		count    int64
	)

	flag = c.FormValue("flag")
	pagesizeStr := c.FormValue("pagesize")
	pageStr := c.FormValue("page")

	if flag == "" {
		return c.JSON(model.OutputRet(model.Notdefinitionparams))
	}
	content = c.FormValue("content")
	if flag == "" {
		return c.JSON(model.OutputRet(model.Notdefinitionparams))
	}

	levelStr := c.FormValue("level")
	if levelStr != "" {
		if level, err = strconv.ParseInt(levelStr, 10, 64); err != nil {
			return c.JSON(model.OutputRet(model.Notdefinitionparams))
		}

		if level != model.LevelLight && level != model.LevelSevere {
			return c.JSON(model.OutputRet(model.Notdefinitionparams))
		}
	} else {
		level = 1
	}

	stateStr := c.FormValue("state")
	if stateStr != "" {

		if state, err = strconv.ParseInt(stateStr, 10, 64); err != nil {
			return c.JSON(model.OutputRet(model.Notdefinitionparams))
		}

		if state != model.StateOpen && state != model.StateClose {
			return c.JSON(model.OutputRet(model.Notdefinitionparams))
		}

	} else {
		state = 1
	}

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil || page < 1 {
		page = 1
	}
	pagesize, err = strconv.ParseInt(pagesizeStr, 10, 64)
	if err != nil || pagesize < 0 || pagesize > 200 {
		pagesize = 100
	}

	if ret, count, err = svr.adminKeywordList(content, flag, state, level, page, pagesize); err != nil {
		return c.JSON(model.OutputRet(model.RetErr))

	}

	if err != nil {
		return c.JSON(model.OutputRet(model.RetErr))
	}

	result := map[string]interface{}{
		"list":  ret,
		"count": count,
	}
	return c.JSON(model.OutputRet(model.Success, result))

}
