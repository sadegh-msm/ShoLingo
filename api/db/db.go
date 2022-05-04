package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var Context = context.Background()

func CreateClients(No int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       No,
	})
	return rdb
}
