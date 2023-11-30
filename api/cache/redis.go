package cache

import (
	"fmt"
	"github.com/go-redis/redis"
)

func NewRedis() *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if client == nil {
		fmt.Println("error could not connect to redis")
	}
	return &RedisCache{
		client: client,
	}
}

type RedisCache struct {
	client *redis.Client
}

func (redis RedisCache) Get() []byte {
	strCmd := redis.client.Get("cities")
	cacheBytes, err := strCmd.Bytes()
	if err != nil {
		return nil
	}
	return cacheBytes
}

func (redis RedisCache) Put(value []byte) {
	redis.client.Set("cities", value, 0)
}
