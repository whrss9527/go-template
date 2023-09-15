package redis

import (
	"context"
	"time"

	"go-template/internal/conf"

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
