package caching

import (
	"time"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/go-redis/redis"
)

// RedisInterface exported
type RedisInterface interface {
	PutOrder(o models.Order) error
	DeleteOrder(merchant string, uuid string) error
	ShowOrder(merchant string, uuid string) (models.Order, error)
	ListOrder(merchant string, uuid string) ([]models.Order, error)
}

//RedisCache exported
type RedisCache struct {
	client *redis.Client
	RedisInterface
}

func (c *RedisCache) getClient() *redis.Client {
	if c.client == nil {
		c.client = redis.NewClient(&redis.Options{
			Addr:        "127.0.0.1:6379",
			Password:    "", // no password set
			DB:          0,  // use default DB
			DialTimeout: 3 * time.Second,
			PoolSize:    300,
		})
	}

	return c.client
}

//Get exported
func (c *RedisCache) Get(key string) (string, error) {
	return c.getClient().Get(key).Result()
}

//HSet exported
func (c *RedisCache) HSet(key, field string, value interface{}) (bool, error) {
	return c.getClient().HSet(key, field, value).Result()
}
