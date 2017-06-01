package usecase

import (
	"encoding/json"
	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

// Persistence exported
type Persistence interface {
	SaveOrder(models.Order) error
	SaveItem(models.Item) error
}

// SaveOrder exported
func (u *UseCase) SaveOrder(o models.Order) error {
	if o.Status == models.CLOSED {
		body, _ := json.Marshal(o)
		dbSave := u.DB.Table("orders").Create(&models.OrderPg{Payload: body})
		return dbSave.Error
	}
	err := cache.PutOrder(o)
	return err
}
