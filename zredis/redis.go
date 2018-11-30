package zredis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Conn struct {
	redis.Conn
}

func dialTLS(addr, password string) (redis.Conn, error) {
	rawURL := fmt.Sprintf("rediss://:%s@%s", password, addr)
	//rawURL := "rediss://:ISEIgnsBspWnLjqPxmpu54DHjYgeVwwSUoTLMkXNnOY=@zhiying-test.redis.cache.chinacloudapi.cn:6380"
	return redis.DialURL(rawURL, redis.DialUseTLS(true), redis.DialTLSSkipVerify(true))
}

func DialTLS(addr, password string) (*Conn, error) {
	// rawURL := fmt.Sprintf("rediss://:%s@%s", password, addr)
	// c, err := redis.DialURL(rawURL, redis.DialUseTLS(true), redis.DialTLSSkipVerify(true))
	c, err := dialTLS(addr, password)
	if nil != err {
		return nil, err
	}
	return &Conn{Conn: c}, nil
}

type Pool struct {
	pool *redis.Pool
}

func NewPoolTLS(addr, password string) *Pool {
	return &Pool{
		pool: &redis.Pool{
			Dial: func() (redis.Conn, error) {
				return DialTLS(addr, password)
			},
		},
	}
}

func (p *Pool) SetMaxIdle(n int) {
	p.pool.MaxIdle = n
}

func (p *Pool) SetMaxActive(n int) {
	p.pool.MaxActive = n
}

func (p *Pool) SetIdleTime(t time.Duration) {
	p.pool.IdleTimeout = t
}

func (p *Pool) Get() *Conn {
	return &Conn{Conn: p.pool.Get()}
}

func (p *Pool) Close() error {
	return p.pool.Close()
}

//for list
func (c *Conn) LrangeStrings(key string, start, to int) ([]string, error) {
	return redis.Strings(c.Do("LRANGE", key, start, to))
}

func (c *Conn) Ltrim(key string, start, to int) (string, error) {
	return redis.String(c.Do("LTRIM", key, start, to))
}
