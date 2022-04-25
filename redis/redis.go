package redis

import (
	"context"
	"errors"
	"fmt"
	"gomod/common"
	"net"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
)

type (
	Endpoint struct {
		Redis []*Redis
	}
	Redis struct {
		Conn    redis.Conn
		tcpConn net.Conn
		ctx     context.Context
		dialer  net.Dialer
		Pool    *redis.Pool
		timeout time.Duration
	}
)

var (
	err error
	r   *Redis
	ep  *Endpoint
)

func New() *Endpoint {
	ep = new(Endpoint)
	return ep
}
func (e *Endpoint) NewRedis() *Redis {
	var r *Redis = NewRedis()
	e.Redis = append(e.Redis, r)
	return r
}
func (e *Endpoint) GetRedis(endpointIndex int) *Redis {
	var len_r = len(e.Redis)
	if len_r < 1 {
		return NewRedis()
	}
	return e.Redis[endpointIndex]
}
func (e *Endpoint) Close() {
	var (
		i, len_r int
	)
	len_r = len(e.Redis)
	for i = 0; i < len_r; i++ {
		e.GetRedis(i).Close()
		e.Redis[i] = nil
	}
	e.Redis = []*Redis{}
}
func NewRedis() *Redis {
	r = new(Redis)
	return r
}
func (r *Redis) Close() {
	if r.Conn != nil {
		r.Conn.Close()
		r.Conn = nil
	}
	if r.Pool != nil {
		r.Pool.Close()
		r.Pool = nil
	}
}
func (r *Redis) SetContext(timeout time.Duration) {
	r.timeout = timeout
	r.ctx, _ = common.Context(timeout)
}
func (r *Redis) NewConn(host string, port int, timeout time.Duration) error {
	r.SetContext(timeout)
	return r.newConnCore(r.ctx, host, port)
}
func (r *Redis) newConnCore(ctx context.Context, host string, port int) error {
	r.tcpConn, err = r.dialer.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", host, port))
	r.Conn = redis.NewConn(r.tcpConn, r.timeout, r.timeout)
	return err
}
func (r *Redis) NewPool(host string, port int, maxActive, maxIdle int, timeout time.Duration) *Redis {
	r.SetContext(timeout)
	r.Pool = &redis.Pool{
		DialContext: func(ctx context.Context) (redis.Conn, error) {
			var err error
			if err = r.newConnCore(ctx, host, port); err != nil {
				return nil, err
			}
			return r.Conn, nil
		},
		IdleTimeout: r.timeout,
		MaxActive:   maxActive,
		MaxIdle:     maxIdle,
	}
	return r
}
func (r *Redis) NewPoolConn() (redis.Conn, error) {
	var conn redis.Conn
	if err != nil {
		return nil, err
	}
	return conn, nil
}
func (r *Redis) GetConn() (redis.Conn, error) {
	if r.Pool == nil {
		if r.Conn == nil {
			return nil, errors.New("no connection")
		} else if ok := r.Ping(); !ok {
			defer r.Conn.Close()
			return nil, errors.New("no connection")
		}
	} else {
		r.Conn = r.Pool.Get()
	}
	return r.Conn, nil
}
func (r *Redis) Do(commandName string, args ...interface{}) (interface{}, error) {
	if _, err = r.GetConn(); err != nil {
		return nil, err
	}
	return r.Conn.Do(commandName, args...)
}
func (r *Redis) Ping() bool {
	var rep string
	rep, err = redis.String(r.Conn.Do("ping"))
	if err != nil {
		return false
	}
	return strings.Compare(rep, "PONG") == 0
}
func (r *Redis) HSetString(key string, keyValue map[string]string) error {
	var (
		err  error
		k, v string
	)

	for k, v = range keyValue {
		_, err = r.Do("hset", key, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
func (r *Redis) HSet(key string, keyValue map[string][]byte) error {
	var (
		err error
		k   string
		v   []byte
	)

	for k, v = range keyValue {
		_, err = r.Do("hset", key, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
func (r *Redis) HGetAll(key string) (map[string][]byte, error) {
	var (
		rep [][]byte
		i   int
		k   string
		v   []byte
		m   = make(map[string][]byte)
	)
	if rep, err = redis.ByteSlices(r.Do("hgetall", key)); err != nil {
		return make(map[string][]byte), err
	}
	for i, v = range rep {
		if common.Number(i).Even() {
			k = string(v)
		} else {
			m[k] = v
		}
	}
	return m, nil
}
func (r *Redis) Keys(keyName string) ([]string, error) {
	var rep []string
	if rep, err = redis.Strings(r.Do("keys", keyName)); err != nil {
		return []string{}, err
	}
	return rep, nil
}
func (r *Redis) Expire(interval int, keys ...string) error {
	var k string
	for _, k = range keys {
		_, err = r.Do("expire", k, interval)
		if err != nil {
			return err
		}
	}
	return nil
}
