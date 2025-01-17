package redis

import (
	"fmt"
	"sync"
	"time"

	"github.com/bwcxyk/golibs/config"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

// https://github.com/go-redis/redis
// 单机模式

var redisOnce sync.Once
var redisNew Client

func NewRedis() Client {
	redisOnce.Do(func() {
		ipAndPort := fmt.Sprintf("%s:%d", config.Config.GetRedis().GetHost(), config.Config.GetRedis().GetPort())
		redisNew = redis.NewClient(&redis.Options{
			Addr:         ipAndPort,
			Username:     config.Config.GetRedis().GetUsername(),
			Password:     config.Config.GetRedis().GetPassword(),
			DB:           config.Config.GetRedis().GetDb(),
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,

			MaxRetries: -1,

			PoolSize:           10,
			PoolTimeout:        30 * time.Second,
			IdleTimeout:        time.Minute,
			IdleCheckFrequency: 100 * time.Millisecond,
		})
		err := redisNew.Ping(context.Background()).Err()
		if err != nil {
			panic(err)
		}
	})
	return redisNew
}
