package identify

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"golang-kit/config"
	"golang-kit/db/mysql"
	"golang-kit/db/redis"
	"golang-kit/ecode"
	"golang-kit/log"
	"golang-kit/net/context"
	xredis "golang/redigo/redis"
	"net/http"
	"news-service/model"
	"strconv"
	"strings"
)

const (
	// cookie
	LOGIN_COOKIE        = "SESSDATA" // the key of login cookie SESSDATA=uid
	LOGIN_COOKIE_EXPIRE = 86400 * 30 // 30 day

	// redis
	REDIS_LOGIN_KEY = "uid_%d" // the value of login cookie key such uid_123
	REDIS_OK        = "OK"
	REDIS_EXPIRE    = 86400 * 30 // 30 day

	// mysql
	_queryUser = "select id from member_base_info where account=? and password=?"

	// 得到用户信息 第一次登陆用
	_getMemberInfoByUidSQL = "select " +
		"mb.id,mb.account,mb.sex,mb.avatar," +
		"ms.level,ms.start_time,ms.end_time," +
		"mc.company_name,mc.business_section,mc.license_pic,mc.contact_name,mc.mobile,mc.phone,mc.qq,mc.email,mc.addr,mc.website,mc.wechat_pic,mc.scope,mc.brief,mc.banner," +
		"msu.name,msu.mobile,msu.phone,msu.wechat,msu.qq,msu.email" +
		" from member_base_info as mb" +
		" inner join member_service_info as ms on mb.id=ms.uid" +
		" inner join member_company_info as mc on mb.id=mc.uid" +
		" inner join member_support_info as msu on mb.id=msu.uid" +
		" where mb.id = ?"

	// 不存在用户
	USER_NOTEXIST int64 = 0

	// context
	CtxMID = "uid"
)

type Service2 struct {
	conf  *config.Identify
	redis *redis.Pool
	mysql *mysql.DB
	keys  map[string]string // key:secret
}

func NewIdentify2(c *config.Identify) (s *Service2, err error) {
	s = &Service2{conf: c}
	s.redis = redis.NewRedisPool(c.Redis)
	s.keys = make(map[string]string)
	if s.mysql, err = mysql.NewMysql(c.Mysql); err != nil {
		log.Error("mysql.NewMysql error(%v)", err)
		return
	}
	return
}

func (s *Service2) Verify(c context.Context) (err error) {
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

func (s *Service2) redisUserKey(uid int64) string {
	return fmt.Sprintf(REDIS_LOGIN_KEY, uid)
}

// 是否登录
func (s *Service2) IsLogin(c context.Context) (err error) {
	// first step set mid = 0
	// set context
	c.Set(CtxMID, 0)

	// get cookie
	cookie, err := c.Request().Cookie(LOGIN_COOKIE)
	if err == http.ErrNoCookie {
		err = ecode.NoLogin
		return
	}
	value := cookie.Value
	var (
		cookieUid int64
		redisUid  int64
	)
	if cookieUid, err = strconv.ParseInt(value, 10, 64); err != nil {
		err = ecode.NoLogin
		return
	}
	conn := s.redis.Get(c)
	defer conn.Close()
	if redisUid, err = xredis.Int64(conn.Do("get", s.redisUserKey(cookieUid))); err != nil {
		log.Error("xredis.Int64(%s) err(%v)", value, err)
		return
	}
	if cookieUid != redisUid {
		log.Error("xredis.Int64(%s) cookieUid(%d), redisUid(%d)", value, cookieUid, redisUid)
		err = ecode.NoLogin
		return
	}

	// set context
	c.Set(CtxMID, redisUid)
	return
}

// 登出
func (s *Service2) Logout(c context.Context) (err error) {
	// get user
	if err = s.IsLogin(c); err != nil {
		return
	}
	var uid int64
	if midInter, ok := c.Get("mid"); ok {
		uid = midInter.(int64)
	}
	// 删除redis
	conn := s.redis.Get(c)
	defer conn.Close()
	if _, err = xredis.Int64(conn.Do("del", s.redisUserKey(uid))); err != nil {
		log.Error("conn.Do(del) uid(%d) err(%v)", uid, err)
		return
	}
	// 删除cookie
	cookie := http.Cookie{Name: LOGIN_COOKIE, Path: "/", MaxAge: -1}
	http.SetCookie(c.Response(), &cookie)
	return
}

// 登录
func (s *Service2) Login(c context.Context) (err error) {
	// 判断是否已经登录
	err = s.IsLogin(c)
	if err == nil {
		err = ecode.Logined
		return
	}
	if err != ecode.NoLogin {

		return
	}
	var uid int64
	// get and checkout user
	if uid, err = s.checkUser(c); err != nil {
		return
	}

	// set redis
	conn := s.redis.Get(c)
	defer conn.Close()
	if err = conn.Send("set", s.redisUserKey(uid), uid); err != nil {
		log.Error("conn.Send(set) key(%s) error(%v)", s.redisUserKey(uid), err)
		return
	}
	if err = conn.Send("expire", s.redisUserKey(uid), REDIS_EXPIRE); err != nil {
		log.Error("conn.Send(set) key(%s) error(%v)", s.redisUserKey(uid), err)
		return
	}
	if err = conn.Flush(); err != nil {
		log.Error("conn.Flush() key(%s) error(%v)", s.redisUserKey(uid), err)
		return
	}
	for i := 0; i < 2; i++ {
		if _, err = conn.Receive(); err != nil {
			log.Error("conn.Do(set) key(%s) err(%v)", s.redisUserKey(uid), err)
			return
		}
	}

	// set cookie
	cookie := http.Cookie{Name: LOGIN_COOKIE, Value: fmt.Sprintf("%d", uid), Path: "/", MaxAge: LOGIN_COOKIE_EXPIRE}
	http.SetCookie(c.Response(), &cookie)

	// set context
	c.Set(CtxMID, uid)

	// 得到用户信息
	var mInfo *model.MemberInfo
	res := c.Result()
	if mInfo, err = s.GetMemberInfoByUid(c, uid); err != nil {
		res["err"] = err
		return
	}
	res["data"] = mInfo
	return
}

func (s *Service2) checkUser(c context.Context) (uid int64, err error) {
	params := c.Request().Form
	account := params.Get("account")
	pass := params.Get("password")
	if uid, err = s.getUserFromDb(c, account, pass); err != nil {
		return
	}
	if uid == USER_NOTEXIST {
		err = ecode.UserNotExist
		return
	}
	return
}

func (s *Service2) getUserFromDb(c context.Context, name, pass string) (uid int64, err error) {
	row := s.mysql.QueryRow(c, _queryUser, name, pass)
	if err = row.Scan(&uid); err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		} else {
			log.Error("s.mysql.QueryRow(%s,%s,%s) error(%v)", _queryUser, name, pass, err)
			return
		}
	}
	return
}

func (s *Service2) GetMemberInfoByUid(c context.Context, uid int64) (mInfo *model.MemberInfo, err error) {
	var (
		mb  = &model.MemberBaseInfo{}
		ms  = &model.MemberServiceInfo{}
		mc  = &model.MemberCompanyInfo{}
		msu = &model.MemberSupportInfo{}
	)
	row := s.mysql.QueryRow(c, _getMemberInfoByUidSQL, uid)

	if err = row.Scan(
		&mb.Id, &mb.Account, &mb.Sex, &mb.Avatar,
		&ms.Level, &ms.StartTime, &ms.EndTime,
		&mc.CompanyName, &mc.BusinessSection, &mc.LicensePic, &mc.ContactName, &mc.Addr, &mc.Mobile, &mc.QQ, &mc.Email, &mc.Addr, &mc.Website, &mc.WechatPic, &mc.Scope, &mc.Brief, &mc.Banner,
		&msu.Name, &msu.Mobile, &msu.Phone, &msu.Wechat, &msu.QQ, &msu.Email); err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		} else {
			log.Error("row.Scan(uid %d) error(%v)", uid, err)
		}
		return
	}
	mInfo = &model.MemberInfo{
		Base:    mb,
		Service: ms,
		Company: mc,
		Support: msu,
	}
	return
}
