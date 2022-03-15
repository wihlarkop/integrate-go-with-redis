package gateway

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

func GetKey(ctx context.Context, redis *redis.Client, prefix string) string {
	val, err := redis.Get(ctx, prefix).Result()
	if err != nil {
		log.Println(err)
	}
	return val
}

func SetKey(ctx context.Context, redis *redis.Client, prefix string, value string) {
	err := redis.Set(ctx, prefix, value, 0).Err()
	if err != nil {
		log.Println(err)
	}
}
