package rediscli

import "github.com/arthurstockler/omaha-order-manager-service-go/models"

// DelOrder in memory
func DelOrder(o models.Order) error {

	_, err := client.delOrder(o.MerchantID, o.UUID.String())

	return err
}
