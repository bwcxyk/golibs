package redis

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/it-sos/golibs/config"
	"golang.org/x/net/context"
)

// https://github.com/go-redis/redis
// 单机模式

type GoLibRedis = *redis.Client

var redisOnce sync.Once
var redisNew GoLibRedis

func NewRedis() GoLibRedis {
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
