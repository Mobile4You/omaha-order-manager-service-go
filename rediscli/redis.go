package rediscli

import "github.com/go-redis/redis"

var redisCli *OrderClient = &OrderClient{}

func init() {

	redisCli.redisClient = redis.NewClient(&redis.Options{
		Addr:     "order-sse-redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}
