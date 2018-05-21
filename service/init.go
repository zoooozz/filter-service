package service

import (
	"context"
	"filter-service/actire"
	"filter-service/model"
	"golang-kit/log"
)

func (s *service) initLoadBusiness() (ts []*model.Business, err error) {

	if ts, err = s.dao.GetByBusinessList(context.TODO()); err != nil {
		panic(err)
	}
	s.businessMap = make(map[string]struct{}, len(ts))
	for _, t := range ts {
		s.businessMap[t.Flag] = struct{}{}
	}
	s.filterMap = make(map[string]*model.Filter, len(ts))
	return
}

func (s *service) initLoadKeyword() {
	filters := make(map[string]*model.Filter, len(s.businessMap))
	for business, _ := range s.businessMap {
		filters[business] = &model.Filter{}
		keywords, err := s.dao.GetByrelationKeyword(context.TODO(), business)
		if err != nil {
			log.Error("s.dao.Rule(%s) err(%v)", business, err)
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
	return

}
