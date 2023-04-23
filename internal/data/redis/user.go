package redis

import (
	"context"
	redisClient "go-template/internal/pkg/redis"
)

func FindShowId(simbaId string) (userId string, err error) {
	var key = "user:" + simbaId
	ctx := context.Background()
	cli := redisClient.RedisClient
	userId, err = cli.Get(ctx, key).Result()
	//id, _ := strconv.ParseInt(res, 10, 64)
	return
}

func SetShowId(simbaId string, userId string) {
	var key = "user:" + simbaId
	ctx := context.Background()
	cli := redisClient.RedisClient
	cli.Set(ctx, key, userId, 0)
}

func UserExist(simbaId string) bool {
	var key = "user:" + simbaId
	ctx := context.Background()
	cli := redisClient.RedisClient
	res, _ := cli.Exists(ctx, key).Result()
	return res != 0
}
