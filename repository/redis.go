package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	RDB *redis.Client
	Ctx = context.Background()
)

func ConnectRedis() error {
	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",

		DB: 0,
	})

	_, err := RDB.Ping(Ctx).Result()
	return err
}

func SetJSON(key string, val []byte, ttl time.Duration) error {
	return RDB.Set(Ctx, key, val, ttl).Err()
}

func GetJSON(key string) ([]byte, error) {
	return RDB.Get(Ctx, key).Bytes()
}

func Del(key string) error {
	return RDB.Del(Ctx, key).Err()
}
