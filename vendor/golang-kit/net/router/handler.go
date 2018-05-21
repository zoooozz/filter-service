package router

import (
	"context"
	"fmt"
	opentracing "github.com/opentracing/opentracing-go"
	"golang-kit/log"
	kitCtx "golang-kit/net/context"
	"net/http"
	"strings"
)

// type Handler interface {
// 	ServeHTTP(context.Context)
// }

type HandlerFunc func(kitCtx.Context)

func (f HandlerFunc) ServeHTTP(c kitCtx.Context) {
	f(c)
}

type Handle struct {
	mux *http.ServeMux
}

func NewHandle(m *http.ServeMux) *Handle {
	return &Handle{mux: m}
}

func (r *Handle) GetFunc(pattern string, handlers ...HandlerFunc) {
	r.mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		handleFunc("GET", w, req, handlers)
	})
	return
}

func (r *Handle) PostFunc(pattern string, handlers ...HandlerFunc) {
	r.mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		handleFunc("POST", w, req, handlers)
	})
	return
}

func handleFunc(method string, w http.ResponseWriter, r *http.Request, handlers []HandlerFunc) {
	log.Info(r.RequestURI)
	if r.Method != method {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	path := fmt.Sprintf("start %s", strings.Split(r.RequestURI, "?")[0])
	span := opentracing.StartSpan(path)
	c := kitCtx.NewContext(opentracing.ContextWithSpan(context.WithValue(context.Background(), "sync", true), span), r, w)
	defer span.Finish()
	defer c.Cancel()
	for _, h := range handlers {
		h(c)
		if err := c.Err(); err != nil {
			return
		}
	}
}
