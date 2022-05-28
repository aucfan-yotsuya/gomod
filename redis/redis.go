package redis

import (
	"bytes"
	"net"
	"time"

	"encoding/gob"

	"github.com/gomodule/redigo/redis"
)

type (
	Redis struct {
		Target  []*Target
		Encoder *gob.Encoder
		Decoder *gob.Decoder
	}
	Target struct {
		Conn         redis.Conn
		RedisConnOpt *RedisConnOpt
		PubSubConn   *redis.PubSubConn
		Pool         *redis.Pool
		netConn      net.Conn
		netDialer    net.Dialer
		Buffer       *bytes.Buffer
		NewEncoder   func() *Target
		NewDecoder   func() *Target
		Encode       func(interface{}) error
		Decode       func(interface{}) error
	}
	RedisConnOpt struct {
		Protocol, Address                      string
		PoolMaxActive, PoolMaxIdle, RetryCount int
		Timeout                                time.Duration
	}
)

var (
	err           error
	r             *Redis
	tg            *Target
	retryCount    = 0
	maxRetryCount = 0
	timeout       time.Duration
)

func New() *Redis {
	return new(Redis)
}
