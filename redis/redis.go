package redis

import (
	"net"
	"time"

	"github.com/gomodule/redigo/redis"
)

type (
	Redis struct {
		Target []*Target
	}
	Target struct {
		Conn      redis.Conn
		Pool      *redis.Pool
		netConn   net.Conn
		netDialer net.Dialer
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
