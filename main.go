package main

import (
	"filter-service/config"
	"filter-service/service"
	"flag"
	"golang-kit/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//confg
	flag.Parse()

	if err := config.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	// log
	log.Init(config.Conf.Log)
	defer log.Close()

	// run server
	log.Info("service start")
	if err := service.Run(config.Conf); err != nil {
		time.Sleep(time.Second)
		log.Info("service fail")
		return
	}
	// exit signal
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT)
	for s := range c {
		switch s {
		case syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT:
			log.Info("service end receive signal %v", s)
			return
		default:
			log.Info("other")
		}
	}
	return
}
