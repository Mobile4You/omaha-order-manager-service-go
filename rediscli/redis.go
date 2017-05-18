package rediscli

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var (
	client OrderClient
)

func init() {

	client.rds = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	client.channels = make(map[string]Channel, 0)

	go func() {
		for v, r := pingRegis(); r != nil; v, r = pingRegis() {
			fmt.Println(v, r)
			time.Sleep(1000 * time.Millisecond)
		}
	}()
}

// executa teste de conexao com Redis
func pingRegis() (string, error) {
	return client.rds.Ping().Result()
}
