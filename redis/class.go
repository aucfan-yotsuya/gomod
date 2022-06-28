package redis

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/aucfan-yotsuya/gomod/common"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) NewTarget() *Target {
	var tg = &Target{
		Encode: r.Encode,
		Decode: r.Decode,
		Buffer: new(bytes.Buffer),
	}
	tg.NewEncoder = func() *Target { r.NewEncoder(tg.Buffer); return tg }
	tg.NewDecoder = func() *Target { r.NewDecoder(tg.Buffer); return tg }
	r.Target = append(r.Target, tg)
	return tg
}
func (r *Redis) NilRedis(index int) bool { return r.Target[index] == nil }
func (r *Redis) TargetLen() int          { return len(r.Target) }
func (r *Redis) IncrRetryCount() bool {
	time.Sleep(timeout)
	if retryCount < maxRetryCount {
		retryCount++
		return true
	}
	return false
}
func (r *Redis) ResetRetryCount() bool {
	retryCount = 0
	return true
}
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
func (tg *Target) NilPool() bool { return tg.Pool == nil }
func (tg *Target) Close() {
	if !tg.NilPool() {
		tg.Pool.Close()
		tg.Pool = nil
	}
}
func (t *Target) IncrRetryCount() bool {
	time.Sleep(timeout)
	if retryCount < maxRetryCount {
		retryCount++
		return true
	}
	return false
}
func (t *Target) ResetRetryCount() bool {
	retryCount = 0
	return true
}
func (tg *Target) NewConn(opt *RedisConnOpt) (*Conn, error) {
	var (
		conn    = new(Conn)
		netConn net.Conn
	)
	tg.RedisConnOpt = opt
	if netConn, err = tg.netDialer.DialContext(
		func() context.Context { ctx, _ := common.Context(opt.Timeout); return ctx }(),
		opt.Protocol,
		opt.Address,
	); err != nil {
		return nil, &Err{Message: err.Error()}
	}
	conn.Conn = redis.NewConn(netConn, opt.Timeout, opt.Timeout)
	if conn == nil {
		return nil, &Err{Message: "redisConn has nil"}
	}
	maxRetryCount = opt.RetryCount
	timeout = opt.Timeout
	return conn, nil
}
func (tg *Target) NewPool(opt *RedisConnOpt) *Conn {
	var conn = new(Conn)
	tg.Pool = &redis.Pool{
		DialContext: func(ctx context.Context) (redis.Conn, error) {
			defer tg.ResetRetryCount()
		retry:
			var (
				conn redis.Conn
				err  error
			)
			if conn, err = tg.NewConn(opt); err != nil {
				if tg.IncrRetryCount() {
					goto retry
				}
				return nil, err
			}
			return conn, nil
		},
		IdleTimeout: opt.Timeout,
		MaxActive:   opt.PoolMaxActive,
		MaxIdle:     opt.PoolMaxIdle,
	}
	conn.Conn = tg.Pool.Get()
	return conn
}
func (tg *Target) NewPubSubConn() *redis.PubSubConn {
	var conn = tg.GetConn()
	tg.PubSubConn = &redis.PubSubConn{
		Conn: conn.Conn,
	}
	return tg.PubSubConn
}
func (tg *Target) GetConn() *Conn {
	var conn = new(Conn)
	if tg.Pool == nil {
		conn, _ = tg.NewConn(tg.RedisConnOpt)
	} else {
		conn.Conn = tg.Pool.Get()
	}
	return conn
}
func (c *Conn) Ping() bool {
	var rep string
	if rep, err = redis.String(c.Do("ping")); err != nil {
		return false
	}
	return strings.Compare(rep, "PONG") == 0
}
func (c *Conn) LpushString(Key string, Value ...string) error {
	if _, err = c.Do("lpush", Key, Value); err != nil {
		return err
	}
	return nil
}
func (c *Conn) Lpush(Key string, Value ...[]byte) error {
	if _, err = c.Do("lpush", Key, Value); err != nil {
		return err
	}
	return nil
}
func (c *Conn) RpushString(Key string, Value ...string) error {
	if _, err = c.Do("rpush", Key, Value); err != nil {
		return err
	}
	return nil
}
func (c *Conn) Rpush(Key string, Value ...[]byte) error {
	if _, err = c.Do("rpush", Key, Value); err != nil {
		return err
	}
	return nil
}
func (c *Conn) SetString(Key, Value string) error {
	if _, err = c.Do("set", Key, Value); err != nil {
		return err
	}
	return nil
}
func (c *Conn) Set(Key string, Value []byte) error {
	if _, err = c.Do("set", Key, Value); err != nil {
		return err
	}
	return nil
}
func (c *Conn) Get(Key string) ([]byte, error) {
	return redis.Bytes(c.Do("get", Key))
}
func (c *Conn) GetString(Key string) (string, error) {
	return redis.String(c.Do("get", Key))
}
func (c *Conn) HSetString(key string, keyValue map[string]string) error {
	for k, v := range keyValue {
		if _, err = c.Do("hset", key, k, v); err != nil {
			return err
		}
	}
	return nil
}
func (c *Conn) HSet(key string, keyValue map[string][]byte) error {
	for k, v := range keyValue {
		if _, err = c.Do("hset", key, k, v); err != nil {
			return err
		}
	}
	return nil
}
func (c *Conn) HGetAll(key string) (map[string][]byte, error) {
	var (
		rep [][]byte
		k   string
		m   = make(map[string][]byte)
	)
	if rep, err = redis.ByteSlices(c.Do("hgetall", key)); err != nil {
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
func (c *Conn) HGetAllString(key string) (map[string]string, error) {
	var (
		rep [][]byte
		k   string
		m   = make(map[string]string)
	)
	if rep, err = redis.ByteSlices(c.Do("hgetall", key)); err != nil {
		return make(map[string]string), err
	}
	for i, v := range rep {
		if common.Number(i).Even() {
			k = string(v)
		} else {
			m[k] = string(v)
		}
	}
	return m, nil
}
func (c *Conn) Lrange(Key string, Value ...string) ([][]byte, error) {
	var resp [][]byte
	if resp, err = redis.ByteSlices(c.Do("lrange", Key, Value)); err != nil {
		return [][]byte{}, err
	}
	return resp, nil
}
func (c *Conn) Llen(Key string) (int, error) {
	var resp int
	if resp, err = redis.Int(c.Do("lrange", Key)); err != nil {
		return 0, err
	}
	return resp, nil
}
func (c *Conn) Dump(Key string) ([]byte, error) {
	var resp []byte
	if resp, err = redis.Bytes(c.Do("dump", Key)); err != nil {
		return []byte{}, err
	}
	return resp, nil
}
func (c *Conn) Restore(Key string, TTL int, Value []byte) error {
	_, err = c.Do("restore", Key, TTL, Value)
	return err
}
func (c *Conn) Keys(keyName string) ([]string, error) {
	var rep []string
	if rep, err = redis.Strings(c.Do("keys", keyName)); err != nil {
		return []string{}, err
	}
	return rep, nil
}
func (c *Conn) GetExpire(key string) (int, error) {
	var resp int
	if resp, err = redis.Int(c.Do("ttl", key)); err != nil {
		return 0, err
	}
	return resp, nil
}
func (c *Conn) SetExpire(key string, interval int) error {
	_, err = c.Do("expire", key, interval)
	return err
}
func (c *Conn) HIncrBy(Key string, Field string, Increment int) (int, error) {
	return redis.Int(c.Do("hincrby", Key, Field, Increment))
}
func (c *Conn) Publish(Channel, Message string) (int64, error) {
	return redis.Int64(c.Do("publish", Channel, Message))
}
func (r *Redis) NewEncoder(buffer *bytes.Buffer) *Redis {
	r.Encoder = gob.NewEncoder(buffer)
	return r
}
func (r *Redis) NewDecoder(buffer *bytes.Buffer) *Redis {
	r.Decoder = gob.NewDecoder(buffer)
	return r
}
func (tg *Target) NewBuffer(data []byte) *Redis {
	tg.Buffer = bytes.NewBuffer(data)
	return r
}
func (r *Redis) Encode(Struct interface{}) error {
	return r.Encoder.Encode(Struct)
}
func (r *Redis) Decode(Struct interface{}) error {
	return r.Decoder.Decode(Struct)
}
func (tg *Target) WriteToBuffer(data []byte) (int, error) {
	return fmt.Fprintln(tg.Buffer, data)
}
func (tg *Target) ReadFromBuffer() []byte {
	return tg.Buffer.Bytes()
}
