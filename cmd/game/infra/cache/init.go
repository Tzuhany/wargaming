package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"wargaming/config"
)

var (
	Rdb *redis.Client
)

func Init() {
	ctx := context.Background()

	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       0,
	})

	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}
