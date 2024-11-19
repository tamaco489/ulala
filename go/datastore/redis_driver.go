package datastore

import (
	"github.com/miyabiii1210/ulala/go/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.EnvConfig.RedisHost,
		Password: "",
	})

	return client
}
