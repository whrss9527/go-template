package redis_sync

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	redisClient "go-template/internal/pkg/redis"
	"math/rand"
	"time"
)

// 重试次数
var retryTimes = 5

// 重试频率
var retryInterval = time.Millisecond * 50

// 锁的默认过期时间
var expiration time.Duration

// GetLock 模拟分布式业务加锁场景
func GetLock(lockK string, lockV string, expiration time.Duration) bool {
	rdb := redisClient.RedisClient
	var ctx, cancel = context.WithCancel(context.Background())
	defer func() {
		// 停止goroutine
		cancel()
	}()
	set, err := rdb.SetNX(ctx, lockK, lockV, expiration).Result()
	if err != nil {
		panic(err.Error())
	}
	// 加锁失败,重试
	return set
}

// FreeLock 释放锁
func FreeLock(ctx context.Context, key string, value interface{}) bool {
	rdb := redisClient.RedisClient
	lua := `
-- 如果当前值与锁值一致,删除key
if redis.call('GET', KEYS[1]) == ARGV[1] then
	return redis.call('DEL', KEYS[1])
else
	return 0
end
`
	scriptKeys := []string{key}

	val, err := rdb.Eval(ctx, lua, scriptKeys, value).Result()
	if err != nil {
		panic(err.Error())
	}

	return val == int64(1)
}

// 生成随机时间
func getRandDuration() time.Duration {
	rand.Seed(time.Now().UnixNano())
	min := 50
	max := 100
	return time.Duration(rand.Intn(max-min)+min) * time.Millisecond
}

// 生成随机值
func getRandValue() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}

// 守护线程
func watchDog(ctx context.Context, rdb *redis.Client, key string, expiration time.Duration, tag string) {
	for {
		select {
		// 业务完成
		case <-ctx.Done():
			fmt.Printf("%s任务完成,关闭%s的自动续期\n", tag, key)
			return
			// 业务未完成
		default:
			// 自动续期
			rdb.PExpire(ctx, key, expiration)
			// 继续等待
			time.Sleep(expiration / 2)
		}
	}
}

// 重试
func retry(ctx context.Context, rdb *redis.Client, key string, value interface{}, expiration time.Duration, tag string) bool {
	i := 1
	for i <= retryTimes {
		fmt.Printf(tag+"第%d次尝试加锁中...\n", i)
		set, err := rdb.SetNX(ctx, key, value, expiration).Result()

		if err != nil {
			panic(err.Error())
		}

		if set == true {
			return true
		}

		time.Sleep(retryInterval)
		i++
	}
	return false
}
