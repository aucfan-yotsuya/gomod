package redis

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
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
func (tg *Target) NewConn(opt *RedisConnOpt) error {
	defer tg.ResetRetryCount()
	tg.RedisConnOpt = opt
retryNetConn:
	if tg.netConn, err = tg.netDialer.DialContext(
		func() context.Context { ctx, _ := common.Context(opt.Timeout); return ctx }(),
		opt.Protocol,
		opt.Address,
	); err != nil {
		if tg.IncrRetryCount() {
			goto retryNetConn
		}
		return &Err{Message: err.Error()}
	}
retryRedisConn:
	tg.Conn = redis.NewConn(tg.netConn, opt.Timeout, opt.Timeout)
	if tg.NilConn() {
		if tg.IncrRetryCount() {
			goto retryRedisConn
		}
		return &Err{Message: "redisConn has nil"}
	}
	maxRetryCount = opt.RetryCount
	timeout = opt.Timeout
	return nil
}
func (tg *Target) NewPool(opt *RedisConnOpt) *Redis {
	tg.Pool = &redis.Pool{
		DialContext: func(ctx context.Context) (redis.Conn, error) {
			defer tg.ResetRetryCount()
		retry:
			var err error
			if err = tg.NewConn(opt); err != nil {
				if tg.IncrRetryCount() {
					goto retry
				}
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
func (tg *Target) NewPubSubConn() *redis.PubSubConn {
	tg.GetConn()
	tg.PubSubConn = &redis.PubSubConn{
		Conn: tg.Conn,
	}
	return tg.PubSubConn
}
func (tg *Target) GetConn() *Target {
	if tg.Pool == nil {
		if tg.Conn == nil {
			tg.NewConn(tg.RedisConnOpt)
		} else if ok := tg.Ping(); !ok {
			tg.NewConn(tg.RedisConnOpt)
		}
	} else {
		tg.Conn = tg.Pool.Get()
	}
	return tg
}
func (tg *Target) Do(commandName string, args ...interface{}) (interface{}, error) {
	if tg.NilConn() {
		tg.GetConn()
	}
	if tg.NilConn() {
		return nil, &Err{Message: "Conn has nil"}
	}
	return tg.Conn.Do(commandName, args...)
}
func (tg *Target) Ping() bool {
	var rep string
	if tg.NilConn() {
		tg.GetConn()
	}
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
func (tg *Target) RpushString(Key string, Value ...string) error {
	if _, err = tg.Do("rpush", Key, Value); err != nil {
		return err
	}
	return nil
}
func (tg *Target) Rpush(Key string, Value ...[]byte) error {
	if _, err = tg.Do("rpush", Key, Value); err != nil {
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
func (tg *Target) Get(Key string) ([]byte, error) {
	return redis.Bytes(tg.Do("get", Key))
}
func (tg *Target) GetString(Key string) (string, error) {
	return redis.String(tg.Do("get", Key))
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
func (tg *Target) HGetAllString(key string) (map[string]string, error) {
	var (
		rep [][]byte
		k   string
		m   = make(map[string]string)
	)
	if rep, err = redis.ByteSlices(tg.Do("hgetall", key)); err != nil {
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
func (tg *Target) Lrange(Key string, Value ...string) ([][]byte, error) {
	var resp [][]byte
	if resp, err = redis.ByteSlices(tg.Do("lrange", Key, Value)); err != nil {
		return [][]byte{}, err
	}
	return resp, nil
}
func (tg *Target) Llen(Key string) (int, error) {
	var resp int
	if resp, err = redis.Int(tg.Do("lrange", Key)); err != nil {
		return 0, err
	}
	return resp, nil
}
func (tg *Target) Dump(Key string) ([]byte, error) {
	var resp []byte
	if resp, err = redis.Bytes(tg.Do("dump", Key)); err != nil {
		return []byte{}, err
	}
	return resp, nil
}
func (tg *Target) Restore(Key string, Value []byte) error {
	_, err = tg.Do("restore", Key, Value)
	return err
}
func (tg *Target) Keys(keyName string) ([]string, error) {
	var rep []string
	if rep, err = redis.Strings(tg.Do("keys", keyName)); err != nil {
		return []string{}, err
	}
	return rep, nil
}
func (tg *Target) GetExpire(key string) (int, error) {
	var resp int
	if resp, err = redis.Int(tg.Do("ttl", key)); err != nil {
		return 0, err
	}
	return resp, nil
}
func (tg *Target) SetExpire(key string, interval int) error {
	_, err = tg.Do("expire", key, interval)
	return err
}
func (tg *Target) HIncrBy(Key string, Field string, Increment int) (int, error) {
	return redis.Int(tg.Do("hincrby", Key, Field, Increment))
}
func (tg *Target) Publish(Channel, Message string) (int64, error) {
	return redis.Int64(tg.Do("publish", Channel, Message))
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
