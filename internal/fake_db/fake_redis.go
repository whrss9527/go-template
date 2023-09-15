package test

import (
	"fmt"

	"github.com/alicebob/miniredis/v2"

	v8 "github.com/go-redis/redis/v8"

	"go-template/internal/pkg/redis"
)

// InitTestRedis 实例化一个可以用于单元测试的redis
func InitTestRedis() {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	// 打开下面命令可以测试链接关闭的情况
	// defer mr.Close()

	redis.RedisClient = v8.NewClient(&v8.Options{
		Addr: mr.Addr(),
	})
	fmt.Println("mini redis addr:", mr.Addr())
}
