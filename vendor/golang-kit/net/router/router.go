package router

import (
	"golang-kit/config"
	"golang-kit/service/identify"
	"net/http"
)

type Router struct {
	c         *config.Router
	Mux       *http.ServeMux
	Hand      *Handle
	Identify  identify.Indentify
	slbSwitch bool
}

func NewRouter(c *config.Router, iden identify.Indentify, mux *http.ServeMux) (r *Router) {
	r = &Router{
		c:        c,
		Mux:      mux,
		Hand:     NewHandle(mux),
		Identify: iden,
	}
	r.slbCheck()
	return
}

// 登录
func (r *Router) Login(p string) {
	r.Hand.PostFunc(p, r.preHandler, r.LoginHandler, r.writerHandler)
}

// 登出
func (r *Router) Logout(p string) {
	r.Hand.GetFunc(p, r.preHandler, r.LogoutHandler, r.writerHandler)
}

func (r *Router) GuestGet(p string, hf HandlerFunc) {
	r.Hand.GetFunc(p, r.preHandler, hf, r.writerHandler)
}

func (r *Router) GuestPost(p string, hf HandlerFunc) {
	r.Hand.PostFunc(p, r.preHandler, hf, r.writerHandler)
}

func (r *Router) UserGet(p string, hf HandlerFunc) {
	r.Hand.GetFunc(p, r.preHandler, r.isLoginHandler, hf, r.writerHandler)
}

func (r *Router) UserPost(p string, hf HandlerFunc) {
	r.Hand.PostFunc(p, r.preHandler, r.isLoginHandler, hf, r.writerHandler)
}

func (r *Router) VerifyGet(p string, hf HandlerFunc) {
	r.Hand.GetFunc(p, r.preHandler, r.identifyHandler, hf, r.writerHandler)
}

func (r *Router) VerifyPost(p string, hf HandlerFunc) {
	r.Hand.PostFunc(p, r.preHandler, r.identifyHandler, hf, r.writerHandler)
}

func (r *Router) slbCheck() {
	r.Hand.GetFunc("/health/check", r.preHandler, r.slbCheckHandler, r.writerHandler)
}

func (r *Router) MonitorPing(hf HandlerFunc) {
	r.Hand.GetFunc("/monitor/ping", hf)
}
