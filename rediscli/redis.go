package rediscli

import (
	"fmt"
	"time"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/go-redis/redis"
)

// OrderMemory exported
type Memory interface {
	PutOrder(o models.Order) error
}

type ORedis struct {
	// Implement the OrderMemory interface.
	Memory
}

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
