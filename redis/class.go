package redis

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/aucfan-yotsuya/gomod/common"

	"github.com/gomodule/redigo/redis"
)

type (
	Endpoint struct {
		Redis []*Redis
	}
	Redis struct {
		Conn    redis.Conn
		Pool    *redis.Pool
		tcpConn net.Conn
		ctx     context.Context
		cf      context.CancelFunc
		dialer  net.Dialer
		timeout time.Duration
	}
	RedisConnOpt struct {
		Context context.Context
		Host    string
		Port    int
		Timeout time.Duration
	}
	RedisPoolOpt struct {
	}
)

func (e *Endpoint) NewRedis() *Redis {
	var r *Redis = NewRedis()
	e.Redis = append(e.Redis, r)
	return r
}
func (e *Endpoint) NilRedis(index int) bool {
	return e.Redis[index] == nil
}
func (e *Endpoint) RedisLen() int {
	return len(e.Redis)
}
func (e *Endpoint) GetRedis(endpointIndex int) *Redis {
	if e.RedisLen() < 1 {
		return nil
	}
	return e.Redis[endpointIndex]
}
func (e *Endpoint) Close() {
	for i := 0; i < e.RedisLen(); i++ {
		e.GetRedis(i).Close()
		e.Redis[i] = nil
	}
	e.Redis = []*Redis{}
}
func NewRedis() *Redis {
	r = new(Redis)
	return r
}
func (r *Redis) NilConn() bool {
	return r.Conn == nil
}
func (r *Redis) NilPool() bool {
	return r.Pool == nil
}
func (r *Redis) Close() {
	if !r.NilConn() {
		r.Conn.Close()
		r.Conn = nil
	}
	if !r.NilPool() {
		r.Pool.Close()
		r.Pool = nil
	}
}
func (r *Redis) NewContext(timeout time.Duration) {
	r.timeout = timeout
	r.ctx, r.cf = common.Context(timeout)
}
func (r *Redis) NewConn(redisConnOpt *RedisConnOpt) error {
	r.tcpConn, err = r.dialer.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", host, port))
	r.Conn = redis.NewConn(r.tcpConn, r.timeout, r.timeout)
	return err
	return r.newConnCore(redisConnOpt)
}
func (r *Redis) newConnCore(ctx context.Context, host string, port int) error {
}
func (r *Redis) NewPool(host string, port int, maxActive, maxIdle int, timeout time.Duration) *Redis {
	var (
		ctx context.Context
		cf  context.CancelFunc
	)
	ctx, cf = r.NewContext(timeout)
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
