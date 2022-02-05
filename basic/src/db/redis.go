package db

import (
	"encoding/json"
	"fmt"
	"juggle/basic/src/lib"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

// Cache 缓存
func Cache(f gin.HandlerFunc, param, key string, empty interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param(param)
		key := fmt.Sprintf(key, id)
		conn := RedisPool.Get()
		defer conn.Close()

		// 从redis取数据
		res, err := redis.Bytes(conn.Do("get", key))
		if err != nil && err == redis.ErrNil {
			f(c)
			result, exists := c.Get("result")
			if !exists {
				result = empty
			}
			data, _ := json.Marshal(result)
			_, err = conn.Do("setex", key, 30, data)
			if err != nil {
				log.Println(err)
			}
			c.JSON(http.StatusOK, result)
		} else {
			json.Unmarshal(res, &empty)
			c.JSON(http.StatusOK, empty)
		}
	}
}
