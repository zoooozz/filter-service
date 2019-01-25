package service

import (
	"filter-service/model"
	"github.com/labstack/echo"
)

func index(c echo.Context) error {
	return nil
}

func list(c echo.Context) error {

	var (
		err   error
		t     string
		k     []string
		level int64
	)

	content := c.FormValue("content")
	if content == "" {
		return c.JSON(model.OutputRet(model.Notdefinitionparams))

	}
	flag := c.FormValue("flag")
	if content == "" {
		return c.JSON(model.OutputRet(model.Notdefinitionparams))
	}
	if t, level, k, err = svr.KeywordFilter(content, flag); err != nil {
		return c.JSON(model.OutputRet(model.RetErr))
	}

	result := map[string]interface{}{
		"content": t,
		"keyword": k,
		"level":   level,
	}
	return c.JSON(model.OutputRet(model.Success, result))
}
