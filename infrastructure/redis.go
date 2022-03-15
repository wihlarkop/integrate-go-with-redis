package infrastructure

import (
	"github.com/go-redis/redis/v8"
	"os"
)

type RedisC struct {
	RedisClient *redis.Client
}

func NewRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return rdb
}
