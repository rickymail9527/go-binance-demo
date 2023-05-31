package redisCtrl

import (
	"fmt"
	"log"
	"testing"

	"github.com/go-redis/redis"
)

var redisKey = "streams=btcusdt@aggTrade"

func Test_TryRedisGet(t *testing.T) {
	RedisNewClient()
	redisClient := GetClient()
	data, err := redisClient.Get(redisKey).Result()
	if err != nil {
		if err == redis.Nil {
			log.Println("redis null, error: ", redisKey, err)
			return
		}
		log.Println("error: ", redisKey, err)
		return
	}
	fmt.Println("redis get: ", data)
}

func Test_TryRedisSet(t *testing.T) {
	RedisNewClient()
	redisClient := GetClient()
	redisClient.Set(redisKey, "123", 0)
}
