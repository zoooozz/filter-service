package service

import (
	"context"
)

func (s *service) KeywordFilter(c context.Context, content, flag string) (t string, level int64, k []string, err error) {

	var (
		_emplyRs = []string{}
	)

	filter, ok := s.filterMap[flag]
	if !ok {
		k = _emplyRs
		t = content
		return
	}
	t, level, k = filter.Matcher.Filter(content)

	if len(k) == 0 {
		k = _emplyRs
	}
	return
}
