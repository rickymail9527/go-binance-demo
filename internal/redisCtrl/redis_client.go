package redisCtrl

import (
	"log"

	"github.com/go-redis/redis"
)

const redisAddr = "localhost:6379"

var client *redis.Client

func RedisNewClient() {
	client = redis.NewClient(&redis.Options{
		Addr:         redisAddr,
		Password:     "",
		DB:           0,
		PoolSize:     30,
		MinIdleConns: 30,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal("Initialize redis: ", pong, err)
	} else {
		log.Println("Initialize redis: ", pong)
	}
}

func GetClient() (c *redis.Client) {
	return client
}
