package caches

import (
	"fmt"
	"sync"
	"time"

	"github.com/bwcxyk/golibs/db/redis"
	"golang.org/x/net/context"
)

// StorageStrategy 策略模式实现的方法
type StorageStrategy interface {
	Set(s string) bool
}

// StorageContext 策略模式上下文
type StorageContext interface {
	SetStrategy(strategy *StorageStrategy)
	Set(v string) bool
}

type storageContext struct {
	storageStrategy StorageStrategy
}

func (s *storageContext) SetStrategy(strategy *StorageStrategy) {
	s.storageStrategy = *strategy
}

func (s *storageContext) Set(v string) bool {
	return s.storageStrategy.Set(v)
}

const (
	signRoot = "sign_%s"
)

type redisDb struct {
	db redis.Client
}

func (r *redisDb) Set(s string) bool {
	k := fmt.Sprintf(signRoot, s)
	is, err := r.db.Exists(context.Background(), k).Result()
	if err != nil {
		panic(err)
	}
	if is > 0 {
		return false
	}
	err = r.db.SetEX(context.Background(), k, 1, time.Minute*5).Err()
	if err != nil {
		panic(err)
	}
	return true
}

var signOnce sync.Once
var signNew StorageContext

// SignSet 验签重放验证
func SignSet(v string) bool {
	signOnce.Do(func() {
		var redisDbs StorageStrategy
		redisDbs = &redisDb{redis.NewClient()}
		signNew = &storageContext{}
		signNew.SetStrategy(&redisDbs)
	})
	return signNew.Set(v)
}
