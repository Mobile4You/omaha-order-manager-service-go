package caching

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

// ShowOrder exported
func (c *RedisCache) ShowOrder(key string, field string) (*models.Order, error) {
	var o models.Order
	jsonOrder, _ := c.getClient().HGet(key, field).Result()
	if len(strings.TrimSpace(jsonOrder)) == 0 {
		return nil, errors.New("order not found")
	}
	err := json.Unmarshal([]byte(jsonOrder), &o)
	return &o, err
}
