package redis

import (
	"github.com/edobtc/cloudkit/config"

	redis "github.com/go-redis/redis/v8"
)

func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Read().RedisHost,
		Password: config.Read().RedisPassword, // defaults to ""
		DB:       config.Read().RedisDB,       // defaults to 0
	})
}
