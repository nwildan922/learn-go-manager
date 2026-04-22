package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "P@ssw0rd123.",
		DB:       0,
	})

	// test connection
	_, err := RDB.Ping(context.Background()).Result()

	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}

	log.Println("Redis connected")
}
