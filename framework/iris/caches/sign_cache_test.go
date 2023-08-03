package caches

import (
	"testing"
	"time"

	"github.com/bwcxyk/golibs/db/redis"
	"github.com/bwcxyk/golibs/utils/random"
	"golang.org/x/net/context"
)

func Test_sign_Set(t *testing.T) {
	s := time.Now().Format("060102150405")
	if !SignSet(s) {
		t.Error("预期不符")
	}
	if SignSet(s) {
		t.Error("预期不符")
	}
	if redis.NewRedisCluster().TTL(context.Background(), "sign_"+s).Val().Seconds() < 1 {
		t.Error("预期不符")
	}
}

func BenchmarkSignSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := time.Now().Format("060102150405") + random.Rand(12, random.RandDigit)
		if !SignSet(s) {
			b.Error("预期不符")
		}
	}
}
