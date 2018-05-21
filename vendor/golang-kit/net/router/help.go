package router

import (
	"encoding/json"
	"golang-kit/ecode"
	"golang-kit/ecode2"
	"golang-kit/log"
	kitCtx "golang-kit/net/context"
	"net/http"
	"strings"
)

var (
	allowOriginHosts = []string{
		".xxxx.com",
	}
)

func (r *Router) preHandler(c kitCtx.Context) {
	req := c.Request()
	resp := c.Response()
	req.ParseForm()
	params := req.Form
	// parse proto.
	pt := _protoJSON
	// set Content-Type
	switch pt {
	case _protoJSON:
		resp.Header().Set("Content-Type", "application/json;charset=utf-8")
	}
	// cross domain
	if params.Get("jsonp") == "jsonp" || params.Get("cross_domain") == "true" {
		origin := req.Header.Get("Origin")
		for _, host := range allowOriginHosts {
			if strings.HasSuffix(origin, host) {
				resp.Header().Set("Access-Control-Allow-Origin", origin)
				resp.Header().Set("Access-Control-Allow-Credentials", "true")
				resp.Header().Set("Access-Control-Allow-Methods", "POST, GET")
				break
			}
		}
	}
	return
}

func (r *Router) LoginHandler(c kitCtx.Context) {
	if err := r.Identify.Login(c); err != nil {
		c.Result()["code"] = err
	} else {
		c.Result()["code"] = ecode.OK
	}
	r.writerHandler(c)
	c.Cancel()
	return
}

func (r *Router) LogoutHandler(c kitCtx.Context) {
	if err := r.Identify.Logout(c); err != nil {
		c.Result()["code"] = err
	} else {
		c.Result()["code"] = ecode.OK
	}
	r.writerHandler(c)
	c.Cancel()
	return
}

func (r *Router) isLoginHandler(c kitCtx.Context) {
	if err := r.Identify.IsLogin(c); err != nil {
		c.Result()["code"] = err
		r.writerHandler(c)
		c.Cancel()
	}
	return
}

func (r *Router) identifyHandler(c kitCtx.Context) {
	if err := r.Identify.Verify(c); err != nil {
		c.Result()["code"] = err
		r.writerHandler(c)
		c.Cancel()
	}
	return
}

func (r *Router) writerHandler(c kitCtx.Context) {
	var (
		bs     []byte
		req    = c.Request()
		resp   = c.Response()
		params = req.Form
		path   = req.URL.Path
		pt     = _protoJSON
	)
	// get proto.
	switch pt {
	case _protoJSON:
		bs = r.JSONResult(c)
	}
	if len(bs) > 0 {
		if _, err := resp.Write(bs); err != nil {
			log.Error("c.Response.Write(%s, %s) failed (%v)", path, params.Encode(), err)
		}
	}
}

func (r *Router) JSONResult(c kitCtx.Context) (bs []byte) {
	var (
		ret = ecode.OK
		res = c.Result()
		err error
	)
	// code
	if res["code"] == nil {
		res["code"] = ret
	} else {
		switch res["code"].(type) {
		case ecode2.Ecode2:
			ec, _ := res["code"].(ecode2.Ecode2)
			res["code"] = ec.ToInt()
			res["message"] = ec.ToString()
		case error:
			ec, _ := res["code"].(error)
			ret = ecode.Lookup(ec)
			res["code"] = int(ret)
			res["message"] = ret.Message()
		default:
			ret = ecode.ServerErr
			res["message"] = ret.Message()
		}
		if ec, ok := res["code"].(error); ok {
			ret = ecode.Lookup(ec)
		} else {
			ret = ecode.ServerErr
		}
	}
	if bs, err = json.Marshal(res); err != nil {
		log.Error("json.Marshal(%v) error(%v)", res, err)
		return
	}
	bs = jsonp(c, bs)
	return
}

func (r *Router) slbCheckHandler(c kitCtx.Context) {
	req := c.Request()
	res := c.Result()
	slbSwitch := req.Form.Get("slb_switch")
	if slbSwitch == "" {
		if r.slbSwitch {
			http.Error(c.Response(), http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			c.Cancel()
		}
		return
	}
	if slbSwitch != "true" && slbSwitch != "false" {
		res["code"] = ecode.RequestErr
		res["message"] = "slb_switch must be true or false"
		return
	}
	r.slbSwitch = slbSwitch == "true"
	res["code"] = ecode.OK
	res["data"] = map[string]interface{}{"slb_switch": r.slbSwitch}
	return
}
