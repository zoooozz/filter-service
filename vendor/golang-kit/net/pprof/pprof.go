package pprof

import (
	"golang-kit/config"
	"golang-kit/log"
	"golang-kit/net/httpsvr"
	xhttp "net/http"
	"net/http/pprof"
)

func Init(c *config.Http) (err error) {
	mux := xhttp.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	if err = httpsvr.RunHttp(c, mux); err != nil {
		log.Error("pprof service error(%v)", err)
	} else {
		log.Info("Run pprof success port:%d", c.Port)
	}
	return
}
