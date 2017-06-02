package usecase

import (
	"encoding/json"
	"time"
	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

// Persistence exported
type Persistence interface {
	Save(models.Order) error
	Update(models.Order) error
	ChangeStatus(models.Order, models.OrderStatus) error
}

// ChangeStatus exported
func (u *Store) ChangeStatus(o models.Order, newStatus models.OrderStatus) error {
	if o.Status == newStatus {
		return nil
	}
	o.Status = newStatus
	o.UpdatedAt = time.Now()

	if newStatus != models.Closed {
		return cache.PutOrder(o)
	}

	return u.moveOrder(o)
}

func (u *Store) moveOrder(o models.Order) error {

	body, _ := json.Marshal(o)
	dbSave := u.Conn.Table("orders").
		Create(&models.OrderPg{
			Payload: body,
			UUID:    o.UUID,
		})

	go cache.DeleteOrder(o.MerchantID, o.UUID)

	return dbSave.Error
}

// Update exported
func (u *Store) Update(o models.Order) error {
	return cache.PutOrder(o)
}

// Save exported
func (u *Store) Save(o models.Order) error {

	if o.Status == models.Closed {
		return u.moveOrder(o)
	}
	err := cache.PutOrder(o)
	return err
}
