package initialize

import (
	"context"
	"project/config"
	"project/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func initRedisClient(redisCfg config.Redis) (redis.UniversalClient, error) {
	var client = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Log.Error("redis connect ping failed, err:", zap.String("name", redisCfg.Name), zap.Error(err))
		return nil, err
	}

	return client, nil
}

func Redis() redis.UniversalClient {
	redisClient, err := initRedisClient(global.Config.Redis)
	if err != nil {
		panic(err)
	}
	return redisClient
}
