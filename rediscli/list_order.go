package rediscli

import (
	"encoding/json"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

// ListOrder returns all orders transaction of an EC
func ListOrder(merchantID string) []models.Order {

	keys, _ := client.rds.HKeys(merchantID).Result()

	orders := make([]models.Order, 0)

	for _, v := range keys {

		var order models.Order

		j, _ := client.rds.HGet(merchantID, v).Result()

		json.Unmarshal([]byte(j), &order)

		orders = append(orders, order)
	}

	return orders
}
