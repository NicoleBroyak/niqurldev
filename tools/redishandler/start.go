package redishandler

import (
	"github.com/go-redis/redis/v8"
)

var Client = Start()
var Ctx = Client.Context()

func Start() *redis.Client {
	Client := redis.NewClient(&redis.Options{
		Addr:     "niqurl-redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return Client
}
