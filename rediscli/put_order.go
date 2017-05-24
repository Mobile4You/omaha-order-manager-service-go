package rediscli

import (
	"encoding/json"
	"errors"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

// PutOrder include transactional order in memory (status DRAFT, ENTERED and PAID)
func (rd *ORedis) PutOrder(o models.Order) error {

	if o.Status == models.CLOSED {
		return errors.New("Not allowed to include order with status equal to closed")
	}

	jsonOrder, _ := json.Marshal(o)

	_, err := client.putOrder(o.MerchantID, o.UUID.Hex(), string(jsonOrder))

	return err
}
