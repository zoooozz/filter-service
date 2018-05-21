package httpsvr

import (
	"fmt"
	"golang-kit/config"
	"golang-kit/log"
	"net"
	"net/http"
	"time"
)

func RunHttp(c *config.Http, mux *http.ServeMux) (err error) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Port))
	if err != nil {
		log.Error("service start fail net.Listen error(%v)", err)
		return
	}
	svr := &http.Server{
		ReadTimeout:  time.Duration(c.ReadTimeout),
		WriteTimeout: time.Duration(c.WriteTimeout),
		Handler:      mux,
	}
	go func() {
		if err = svr.Serve(listen); err != nil {
			log.Error("svr.Serve error(%v)", err)
		}
	}()
	log.Error("service start success")
	return
}
