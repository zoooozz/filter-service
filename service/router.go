package service

import (
	"github.com/labstack/echo"
)

func initRouter(e *echo.Echo) {
	e.GET("/", index)
	outter := e.Group("/x/outter", HeaderVerifier())
	{
		outter.POST("/filter/list", list)
	}

	internal := e.Group("/x/internal", HeaderVerifier())
	{
		internal.POST("/filter/keyword/add", addKeyword)
		internal.POST("/filter/keyword/edit", editKeyword)
		internal.POST("/filter/keyword/del", delKeyword)
		internal.GET("/filter/keyword/list", adminKeywordList)
	}
}
