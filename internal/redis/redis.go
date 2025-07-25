package redis

import (
	"fmt"
	"os"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Init() {

	Client = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB: 0,
	})

}
