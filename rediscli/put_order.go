package rediscli

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/go-redis/redis"
)

// OrderClient is an exported
type OrderClient struct {
	sync.RWMutex
	redisClient *redis.Client
	//channels    map[string]*Channel
}

func (c *OrderClient) memOrder(merchant string, number string, value string) (bool, error) {
	c.Lock()
	defer c.Unlock()
	result := c.redisClient.HSet(merchant, number, value)
	return result.Result()
}

// include transactional order in memory (status DRAFT, ENTERED and PAID)
func PutOrder(o models.Order) error {

	if o.Status == models.CLOSED {
		return errors.New("Not allowed to include order with status equal to closed")
	}

	jsonOrder, _ := json.Marshal(o)

	_, err := redisCli.memOrder(o.MerchantID, o.UUID.String(), string(jsonOrder))

	return err
}
