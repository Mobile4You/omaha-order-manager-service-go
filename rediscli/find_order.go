package rediscli

import "github.com/arthurstockler/omaha-order-manager-service-go/models"

// FindOrder returns an order in memory
func FindOrder(merchantID string, number string) (*models.Order, error) {
	return client.getOrder(merchantID, number)
}
