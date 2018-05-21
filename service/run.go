package service

import (
	"filter-service/config"
	kitCfg "golang-kit/config"
	"golang-kit/log"
	"golang-kit/net/httpsvr"
	"golang-kit/net/pprof"
	"golang-kit/net/router"
	"golang-kit/service/identify"
	"net/http"
)

var (
	// server
	svr *service
)

func Run(c *config.Config) (err error) {

	if svr, err = NewService(config.Conf); err != nil {
		return
	}

	// run pprof
	if err = pprof.Init(c.Mhttp.Pprof); err != nil {
		return
	}

	// run http
	if err = runHttp(c.Mhttp, c.Router); err != nil {
		return
	}
	return
}

// http
func runHttp(c *kitCfg.Mhttp, cr *kitCfg.Router) (err error) {
	// internal
	inMux := http.NewServeMux()
	iden := identify.NewIdentify(cr.Indentify)
	if err != nil {
		log.Error("identify.NewIdentify2 error(%v)", err)
		return
	}

	inRou := router.NewRouter(cr, iden, inMux)
	initInner(inRou)
	if err = httpsvr.RunHttp(c.Inner, inMux); err != nil {
		log.Error("httpsvr.RunHttp error(%v)", err)
		return
	}

	// outter
	outMux := http.NewServeMux()
	outRou := router.NewRouter(cr, iden, outMux)
	initOutter(outRou)
	if err = httpsvr.RunHttp(c.Outter, outMux); err != nil {
		log.Error("RunOutterHttp error(%v)", err)
		return
	} else {
		log.Info("RunOutterHttp success port:%d", c.Outter.Port)
	}
	return
}
