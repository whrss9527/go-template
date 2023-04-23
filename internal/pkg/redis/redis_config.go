package redis

import (
	"context"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"go-template/internal/conf"
	"time"

	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
)

// RedisClient redis 客户端
var RedisClient *redis.Client

// ErrRedisNotFound not exist in redis
const ErrRedisNotFound = redis.Nil

// Config redis config
type Config struct {
	Addr         string
	Password     string
	DB           int
	MinIdleConn  int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	PoolTimeout  time.Duration
	// tracing switch
	EnableTrace bool
}

// Init 实例化一个redis client
func Init(c *conf.Data) *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.DB),
		MinIdleConns: int(c.Redis.MinIdleConn),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		PoolSize:     int(c.Redis.PoolSize),
		PoolTimeout:  c.Redis.PoolTimeout.AsDuration(),
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	// hook tracing (using open telemetry)
	if c.Redis.IsTrace {
		RedisClient.AddHook(redisotel.NewTracingHook())
	}

	return RedisClient
}

// InitTestRedis 实例化一个可以用于单元测试的redis
func InitTestRedis() {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	// 打开下面命令可以测试链接关闭的情况
	// defer mr.Close()

	RedisClient = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	fmt.Println("mini redis addr:", mr.Addr())
}
