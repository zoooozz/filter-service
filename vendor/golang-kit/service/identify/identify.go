package identify

import (
	"crypto/md5"
	"encoding/hex"
	"golang-kit/config"
	"golang-kit/ecode"
	"golang-kit/log"
	"golang-kit/net/context"
	"strings"
)

type Service struct {
	c    *config.Identify
	keys map[string]string // key:secret
}

func NewIdentify(c *config.Identify) *Service {
	return &Service{
		c: c,
		keys: map[string]string{
			c.App.Key: c.App.Secret,
		},
	}
}

func (s *Service) Verify(c context.Context) (err error) {
	var (
		r      = c.Request()
		query  = r.Form
		secret string
	)
	if r.Method == "POST" {
		p := r.URL.Query()
		if p.Get("sign") != "" {
			query = p
		}
	}
	tsStr := query.Get("ts")
	if tsStr == "" {
		log.Error("ts(%s) is empty", tsStr)
		err = ecode.RequestErr
		return
	}
	sign := query.Get("sign")
	query.Del("sign")
	appkey := query.Get("appkey")
	if v, ok := s.keys[appkey]; ok {
		secret = v
	} else {
		log.Error("appkey(%s)", appkey)
		err = ecode.SignCheckErr
		return
	}
	tmp := query.Encode()
	if strings.IndexByte(tmp, '+') > -1 {
		tmp = strings.Replace(tmp, "+", "%20", -1)
	}
	mh := md5.Sum([]byte(strings.ToLower(tmp)))
	if hex.EncodeToString(mh[:]) != sign {
		mh1 := md5.Sum([]byte(tmp + secret))
		if hex.EncodeToString(mh1[:]) != sign {
			log.Error("Get sign: %s, expect %x", sign, mh1)
			err = ecode.SignCheckErr

		}
	}
	return
}

func (s *Service) Login(c context.Context) (err error) {
	return
}

func (s *Service) Logout(c context.Context) (err error) {
	return
}

func (s *Service) IsLogin(c context.Context) (err error) {
	return
}
