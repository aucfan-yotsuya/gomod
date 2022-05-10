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
		Conn    redis.Conn
		Pool    *redis.Pool
		tcpConn net.Conn
		dialer  net.Dialer
	}
	RedisConnOpt struct {
		Host                             string
		Port, PoolMaxActive, PoolMaxIdle int
		Timeout                          time.Duration
	}
)

var (
	err error
	r   *Redis
	tg  *Target
)

func New() *Target {
	return new(Target)
}
