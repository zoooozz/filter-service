package config

import (
	kitTime "golang-kit/time"
)

// Common
type Common struct {
	Version    string // version
	Env        string // environment
	Family     string // Family
	HostPort   string // addr
	ServerName string //
}

// log
type Log struct {
	Dir string
}

// trace
type Trace struct {
	Addr          string
	Debug         bool
	SameSpan      bool
	TraceID128Bit bool
}

// mysql
type Mysql struct {
	Addr   string
	DSN    string
	Active int
	Idle   int
}

// redis
type Redis struct {
	Name         string
	Proto        string
	Addr         string
	Auth         string
	Active       int
	Idle         int
	DialTimeout  kitTime.Duration
	ReadTimeout  kitTime.Duration
	WriteTimeout kitTime.Duration
	IdleTimeout  kitTime.Duration
}

// memcache
type Memcache struct {
	Name         string
	Proto        string
	Addr         string
	Auth         string
	Active       int
	Idle         int
	DialTimeout  kitTime.Duration
	ReadTimeout  kitTime.Duration
	WriteTimeout kitTime.Duration
	IdleTimeout  kitTime.Duration
}

// grpc server
type GrpcServer struct {
	Name         string
	Addr         string
	Port         int
	RegisterAddr string
}

// grpc client
type GrpcClient struct {
	ServiceName  string
	RegisterAddr string
}

// =================================== HTTP ==================================
// http server
type Mhttp struct {
	Inner  *Http
	Outter *Http
	Pprof  *Http
}

type Http struct {
	Port         int
	ReadTimeout  kitTime.Duration
	WriteTimeout kitTime.Duration
}

// http client
type HTTPClient struct {
	Dial      kitTime.Duration
	Timeout   kitTime.Duration
	KeepAlive kitTime.Duration
}

// =================================== HTTP ==================================

type Pprof struct {
	Port int
}

type Router struct {
	Indentify *Identify
}

type Identify struct {
	App   *App
	Redis *Redis // can be nil
	Mysql *Mysql // can be nil
}

type App struct {
	Key    string
	Secret string
}
