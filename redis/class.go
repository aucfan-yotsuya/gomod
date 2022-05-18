package redis

import (
	"context"
	"strings"

	"github.com/aucfan-yotsuya/gomod/common"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) NewTarget() *Target {
	var tg = new(Target)
	r.Target = append(r.Target, tg)
	return tg
}
func (r *Redis) NilRedis(index int) bool { return r.Target[index] == nil }
func (r *Redis) TargetLen() int          { return len(r.Target) }
func (r *Redis) GetTarget(index int) *Target {
	if r.TargetLen() < 1 {
		return nil
	}
	return r.Target[index]
}
func (r *Redis) Close() {
	for i := 0; i < r.TargetLen(); i++ {
		r.GetTarget(i).Close()
		r.Target[i] = nil
	}
	r.Target = []*Target{}
}
func (tg *Target) NilConn() bool { return tg.Conn == nil }
func (tg *Target) NilPool() bool { return tg.Pool == nil }
func (tg *Target) Close() {
	if !tg.NilConn() {
		tg.Conn.Close()
		tg.Conn = nil
	}
	if !tg.NilPool() {
		tg.Pool.Close()
		tg.Pool = nil
	}
}
func (tg *Target) NewConn(opt *RedisConnOpt) error {
	if tg.tcpConn, err = tg.dialer.DialContext(
		func() context.Context { ctx, _ := common.Context(opt.Timeout); return ctx }(),
		opt.Protocol,
		opt.Address,
	); err != nil {
		return &Err{Message: err.Error()}
	}
	tg.Conn = redis.NewConn(tg.tcpConn, opt.Timeout, opt.Timeout)
	if tg.NilConn() {
		return &Err{Message: "redisConn has nil"}
	}
	return nil
}
func (tg *Target) NewPool(opt *RedisConnOpt) *Redis {
	tg.Pool = &redis.Pool{
		DialContext: func(ctx context.Context) (redis.Conn, error) {
			var err error
			if err = tg.NewConn(opt); err != nil {
				return nil, err
			}
			return tg.Conn, nil
		},
		IdleTimeout: opt.Timeout,
		MaxActive:   opt.PoolMaxActive,
		MaxIdle:     opt.PoolMaxIdle,
	}
	return r
}
func (tg *Target) GetConn() (redis.Conn, error) {
	if tg.Pool == nil {
		if tg.Conn == nil {
			return nil, &Err{Message: "no connection"}
		} else if ok := tg.Ping(); !ok {
			defer tg.Conn.Close()
			return nil, &Err{Message: "no connection"}
		}
	} else {
		tg.Conn = tg.Pool.Get()
	}
	return tg.Conn, nil
}
func (tg *Target) Do(commandName string, args ...interface{}) (interface{}, error) {
	if _, err = tg.GetConn(); err != nil {
		return nil, err
	}
	return tg.Conn.Do(commandName, args...)
}
func (tg *Target) Ping() bool {
	var rep string
	if rep, err = redis.String(tg.Conn.Do("ping")); err != nil {
		return false
	}
	return strings.Compare(rep, "PONG") == 0
}
func (tg *Target) LpushString(Key string, Value ...string) error {
	if _, err = tg.Do("lpush", Key, Value); err != nil {
		return err
	}
	return nil
}
func (tg *Target) Lpush(Key string, Value ...[]byte) error {
	if _, err = tg.Do("lpush", Key, Value); err != nil {
		return err
	}
	return nil
}
func (tg *Target) SetString(Key, Value string) error {
	if _, err = tg.Do("set", Key, Value); err != nil {
		return err
	}
	return nil
}
func (tg *Target) Set(Key string, Value []byte) error {
	if _, err = tg.Do("set", Key, Value); err != nil {
		return err
	}
	return nil
}
func (tg *Target) HSetString(key string, keyValue map[string]string) error {
	for k, v := range keyValue {
		if _, err = tg.Do("hset", key, k, v); err != nil {
			return err
		}
	}
	return nil
}
func (tg *Target) HSet(key string, keyValue map[string][]byte) error {
	for k, v := range keyValue {
		if _, err = tg.Do("hset", key, k, v); err != nil {
			return err
		}
	}
	return nil
}
func (tg *Target) HGetAll(key string) (map[string][]byte, error) {
	var (
		rep [][]byte
		k   string
		m   = make(map[string][]byte)
	)
	if rep, err = redis.ByteSlices(tg.Do("hgetall", key)); err != nil {
		return make(map[string][]byte), err
	}
	for i, v := range rep {
		if common.Number(i).Even() {
			k = string(v)
		} else {
			m[k] = v
		}
	}
	return m, nil
}
func (tg *Target) Keys(keyName string) ([]string, error) {
	var rep []string
	if rep, err = redis.Strings(tg.Do("keys", keyName)); err != nil {
		return []string{}, err
	}
	return rep, nil
}
func (tg *Target) Expire(interval int, keys ...string) error {
	var k string
	for _, k = range keys {
		_, err = tg.Do("expire", k, interval)
		if err != nil {
			return err
		}
	}
	return nil
}
