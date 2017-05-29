package caching

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

// PutOrder include transactional order in memory (status DRAFT, ENTERED and PAID)
func (c *RedisCache) PutOrder(o models.Order) error {

	if o.Status == models.CLOSED {
		return errors.New("Not allowed to include order with status equal to closed")
	}

	body, _ := json.Marshal(o)

	fmt.Printf("Cliente %v", c.client)
	_, err := c.getClient().HSet(o.MerchantID, o.UUID.Hex(), string(body)).Result()

	return err
}
