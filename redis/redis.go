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
		Protocol, Address          string
		PoolMaxActive, PoolMaxIdle int
		Timeout                    time.Duration
	}
)

var (
	err error
	r   *Redis
	tg  *Target
)

func New() *Redis {
	return new(Redis)
}
