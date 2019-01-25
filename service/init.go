package service

import (
	"filter-service/actire"
	"filter-service/model"
	"github.com/donnie4w/go-logger/logger"
)

func (s *service) initLoadBusiness() (rs []*model.Business, err error) {

	if rs, err = s.dao.GetByBusinessList(); err != nil {
		panic(err)
	}
	s.businessMap = make(map[string]struct{}, len(rs))
	for _, t := range rs {
		s.businessMap[t.Flag] = struct{}{}
	}
	s.filterMap = make(map[string]*model.Filter, len(rs))
	return
}

func (s *service) initLoadKeyword() {

	filters := make(map[string]*model.Filter, len(s.businessMap))
	for business, _ := range s.businessMap {
		filters[business] = &model.Filter{}
		keywords, err := s.dao.GetByrelationKeyword(business)
		if err != nil {
			return
		}
		var (
			matcher = actire.NewMatcher()
		)
		for _, keyword := range keywords {
			matcher.Insert(keyword.Content, keyword.Id, keyword.Level)
		}
		matcher.Build()
		filters[business].Matcher = matcher
	}
	s.filterMap = filters
	logger.Info("重新更新敏感词")
	return

}
