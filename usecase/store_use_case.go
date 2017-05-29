package usecase

import (
	"log"

	"github.com/arthurstockler/omaha-order-manager-service-go/db"
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
		db := db.MgoDb{}
		db.Open()
		err := db.Db.C("order").Insert(&o)
		db.Close()
		return err
	}
	err := cache.PutOrder(o)
	return err
}

// SaveItem exported
func (u *UseCase) SaveItem(i models.Item) error {
	log.Print("metodo real SaveItem")
	return nil
}
