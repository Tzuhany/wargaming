package cache

import (
	"common/config"
	"context"
	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client
)

func Init() {
	ctx := context.Background()

	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Addr,
		Password: config.Conf.Redis.Password,
		DB:       0,
	})

	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}
