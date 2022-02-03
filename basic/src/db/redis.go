package db

import (
	"fmt"
	"juggle/basic/src/lib"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	// RedisPool redis连接池
	RedisPool *redis.Pool
)

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     2,
		IdleTimeout: time.Second * 60,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}

			if _, err := conn.Do("auth", lib.Config.Redis.Password); err != nil {
				conn.Close()
				return nil, err
			}

			return conn, nil
		},
	}
}

func InitRedis() {
	RedisPool = newPool(fmt.Sprintf("%s:%d", lib.Config.Redis.Host, lib.Config.Redis.Port))
}
