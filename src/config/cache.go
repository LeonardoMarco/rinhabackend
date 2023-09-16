package config

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("RINHA_CACHE_URL"),
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()

	if err != nil {
		fmt.Println("Not connect for cache application")
	}

	if err != nil {
		fmt.Println(err)
	}

	return client
}
