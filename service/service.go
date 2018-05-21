package service

import (
	"filter-service/config"
	"filter-service/dao"
	"filter-service/model"
	"time"
)

type service struct {
	dao         *dao.Dao
	businessMap map[string]struct{}
	filterMap   map[string]*model.Filter
}

func NewService(c *config.Config) (s *service, err error) {
	// 拿到dao
	s = &service{}
	if s.dao, err = dao.NewDao(c); err != nil {
		return
	}
	//初始化map
	s.initLoadBusiness()
	s.initLoadKeyword()
	go s.loadKeyword()
	return
}

func (s *service) loadKeyword() {
	for {
		time.Sleep(time.Duration(config.Conf.Time.Tick))
		s.initLoadKeyword()
	}
}
