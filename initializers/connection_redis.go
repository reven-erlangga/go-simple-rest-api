package initializers

import (
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func ConnectRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})

	RDB = rdb

	print("Starting redis server")
}