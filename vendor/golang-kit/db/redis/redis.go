package redis

import (
	"context"
	"fmt"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"golang-kit/config"
	"golang/redigo/redis"
)

type conn struct {
	p   *Pool
	c   redis.Conn
	ctx context.Context
}

// Pool redis pool.
type Pool struct {
	*redis.Pool
	c *config.Redis
}

// NewPool new a redis pool.
func NewRedisPool(c *config.Redis) (p *Pool) {
	p = &Pool{c: c}
	cnop := redis.DialConnectTimeout(time.Duration(c.DialTimeout))
	rdop := redis.DialReadTimeout(time.Duration(c.ReadTimeout))
	wrop := redis.DialWriteTimeout(time.Duration(c.WriteTimeout))
	auop := redis.DialPassword(c.Auth)
	// new pool
	p.Pool = redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial(c.Proto, c.Addr, cnop, rdop, wrop, auop)
	}, c.Idle)
	p.IdleTimeout = time.Duration(c.IdleTimeout)
	p.MaxActive = c.Active
	return
}

func NewConn(c *config.Redis) (cn redis.Conn, err error) {
	cnop := redis.DialConnectTimeout(time.Duration(c.DialTimeout))
	rdop := redis.DialReadTimeout(time.Duration(c.ReadTimeout))
	wrop := redis.DialWriteTimeout(time.Duration(c.WriteTimeout))
	auop := redis.DialPassword(c.Auth)
	// new conn
	cn, err = redis.Dial(c.Proto, c.Addr, cnop, rdop, wrop, auop)
	return
}

func (p *Pool) Get(c context.Context) redis.Conn {
	return &conn{p: p, c: p.Pool.Get(), ctx: c}
}

func (p *Pool) Close() error {
	return p.Pool.Close()
}

func (c *conn) Err() error {
	return c.c.Err()
}

func (c *conn) Close() error {
	return c.c.Close()
}

func key(args interface{}) (key string) {
	keys, _ := args.([]interface{})
	if keys != nil && len(keys) > 0 {
		key, _ = keys[0].(string)
	}
	return
}

func (c *conn) Do(cmd string, args ...interface{}) (reply interface{}, err error) {
	if c.ctx.Value("sync") == true {
		span, _ := opentracing.StartSpanFromContext(c.ctx, fmt.Sprintf("redis %s", cmd))
		span.LogEvent(fmt.Sprintf("%s %v", cmd, args))
		defer span.Finish()
	}
	reply, err = c.c.Do(cmd, args...)
	return
}

func (c *conn) Send(cmd string, args ...interface{}) (err error) {
	if c.ctx.Value("sync") == true {
		span, _ := opentracing.StartSpanFromContext(c.ctx, fmt.Sprintf("redis %s", cmd))
		span.LogEvent(fmt.Sprintf("%s %v", cmd, args))
		defer span.Finish()
	}
	err = c.c.Send(cmd, args...)
	return
}

func (c *conn) Flush() error {
	return c.c.Flush()
}

func (c *conn) Receive() (reply interface{}, err error) {
	if c.ctx.Value("sync") == true {
		span, _ := opentracing.StartSpanFromContext(c.ctx, "redis receive")
		defer span.Finish()
	}
	reply, err = c.c.Receive()
	return
}
