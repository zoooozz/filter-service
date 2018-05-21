package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"golang-kit/config"
	"golang-kit/time"
)

type Config struct {
	Common *config.Common
	Log    *config.Log
	Router *config.Router
	Mysql  *Mysql
	Mhttp  *config.Mhttp
	Time   *Time
}

type Time struct {
	Tick time.Duration
}

type Mysql struct {
	Master *config.Mysql
}

var (
	Conf     = &Config{}
	ConfPath string
)

func init() {
	flag.StringVar(&ConfPath, "conf", "filter.toml", "config path")
}

func Init() (err error) {
	_, err = toml.DecodeFile(ConfPath, &Conf)
	return
}
