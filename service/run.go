package service

import (
	"filter-service/config"
	"filter-service/dao"
	"filter-service/model"
	"fmt"
	"github.com/donnie4w/go-logger/logger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"time"
)

var (
	svr *service
)

type service struct {
	dao         *dao.Dao
	businessMap map[string]struct{}
	filterMap   map[string]*model.Filter
}

func Run() (err error) {
	e := echo.New()
	e.Logger.SetLevel(99)
	e.Use(middleware.Recover())
	if svr, err = initService(); err != nil {
		panic(err)
	}
	initRouter(e)
	initLog()
	e.Start(":" + fmt.Sprintf("%d", config.Conf.Http.Port))
	return
}

func initLog() {
	logger.SetConsole(false)
	logger.SetRollingDaily(config.Conf.Log.Addr, config.Conf.Log.Dir)
	logger.SetLevel(logger.INFO)
}

func initService() (s *service, err error) {
	s = &service{}
	if s.dao, err = dao.NewDao(config.Conf); err != nil {
		return
	}
	//加载项在这边
	s.initLoadBusiness()
	s.initLoadKeyword()
	go s.loadKeyword()
	return
}

func (s *service) loadKeyword() {
	for {
		time.Sleep(5000 * time.Millisecond)

		s.initLoadKeyword()
	}
}
