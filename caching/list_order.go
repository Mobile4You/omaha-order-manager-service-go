package caching

import (
	"encoding/json"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

// ListOrder returns all orders transaction of an EC
func (c *RedisCache) ListOrder(merchantID string) []models.Order {

	keys, _ := c.getClient().HKeys(merchantID).Result()

	orders := make([]models.Order, 0)

	for _, v := range keys {

		var order models.Order

		j, _ := c.getClient().HGet(merchantID, v).Result()

		json.Unmarshal([]byte(j), &order)

		orders = append(orders, order)
	}

	return orders
}
