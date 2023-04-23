package redis_sync

import (
	"context"
	redisClient "go-template/internal/pkg/redis"
	"testing"
	"time"
)

func TestRedisLock(t *testing.T) {
	ctx := context.Background()
	redisClient.InitTestRedis()
	r1 := GetLock("11", "11", time.Duration(5)*time.Second)
	//r2 := GetLock("11", "11", time.Duration(5)*time.Second)

	println(r1)
	//println(r2)

	time.Sleep(10 * time.Second)
	FreeLock(ctx, "11", "11")
	r2 := GetLock("11", "11", time.Duration(1)*time.Second)
	println(r2)
}
