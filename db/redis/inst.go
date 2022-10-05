package redis

import (
	"io"

	"github.com/go-redis/redis/v8"
	"github.com/it-sos/golibs/config"
	"github.com/it-sos/golibs/global/consts"
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
