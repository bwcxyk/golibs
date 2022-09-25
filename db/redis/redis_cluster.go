package redis

import (
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/it-sos/golibs/config"
	"golang.org/x/net/context"
)

// https://github.com/go-redis/redis
// 集群模式

type GoLibRedisCluster = *redis.ClusterClient

var redisClusterOnce sync.Once
var redisClusterNew GoLibRedisCluster

func NewRedisCluster() GoLibRedisCluster {
	redisClusterOnce.Do(func() {
		redisClusterNew = redis.NewClusterClient(&redis.ClusterOptions{
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,

			MaxRedirects: 8,

			PoolSize:           10,
			PoolTimeout:        30 * time.Second,
			IdleTimeout:        time.Minute,
			IdleCheckFrequency: 100 * time.Millisecond,
			Addrs:              config.GetRedisCluster().GetHosts(),

			Username: config.GetRedisCluster().GetUsername(),
			Password: config.GetRedisCluster().GetPassword(),
		})
		err := redisClusterNew.Ping(context.Background()).Err()
		if err != nil {
			panic(err)
		}
	})
	return redisClusterNew
}
