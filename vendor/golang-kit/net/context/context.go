package context

import (
	ctx "context"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Context web context interface
type Context interface {
	ctx.Context
	Request() *http.Request
	Response() http.ResponseWriter
	Result() map[string]interface{}
	Cancel()
	Now() time.Time
	Get(string) (interface{}, bool)
	Set(string, interface{})
	RemoteIP() string
}

type webCtx struct {
	ctx.Context
	cancel   ctx.CancelFunc
	req      *http.Request
	resp     http.ResponseWriter
	res      map[string]interface{}
	now      time.Time
	lock     sync.RWMutex
	data     map[string]interface{}
	remoteIP string
}

func NewContext(c ctx.Context, req *http.Request, resp http.ResponseWriter) Context {
	wc := &webCtx{req: req, resp: resp, now: time.Now()}
	wc.Context, wc.cancel = ctx.WithCancel(c)
	wc.remoteIP = remoteIP(req)
	return wc
}

func (c *webCtx) Request() *http.Request {
	return c.req
}

func (c *webCtx) Response() http.ResponseWriter {
	return c.resp
}

func (c *webCtx) Cancel() {
	c.cancel()
}

func (c *webCtx) Now() time.Time {
	return c.now
}

func (c *webCtx) Result() (res map[string]interface{}) {
	if res = c.res; res == nil {
		res = make(map[string]interface{})
		c.res = res
	}
	return
}

func (c *webCtx) Get(key string) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if c.data == nil {
		return nil, false
	}
	v, ok := c.data[key]
	return v, ok
}

func (c *webCtx) Set(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.data == nil {
		c.data = make(map[string]interface{})
	}
	c.data[key] = value
	return
}

func (c *webCtx) RemoteIP() string {
	return c.remoteIP
}

func remoteIP(r *http.Request) (remote string) {
	var xff = r.Header.Get("X-Forwarded-For")
	if idx := strings.IndexByte(xff, ','); idx > -1 {
		if remote = strings.TrimSpace(xff[:idx]); remote != "" {
			return
		}
	}
	if remote = r.Header.Get("X-Real-IP"); remote != "" {
		return
	}
	remote = r.RemoteAddr[:strings.Index(r.RemoteAddr, ":")]
	return
}
