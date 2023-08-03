package redis

import (
	"io"

	"github.com/bwcxyk/golibs/config"
	"github.com/bwcxyk/golibs/global/consts"
	"github.com/go-redis/redis/v8"
)

// Client redis interface
type Client interface {
	redis.Cmdable // Commands.
	io.Closer     // CloseConnection.
}

// NewClient redis
func NewClient() Client {
	if config.Config.GetRedisUse() == consts.TypeRedis {
		return NewRedis()
	}
	return NewRedisCluster()
}
