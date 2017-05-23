package usecase

import (
	"log"

	"github.com/arthurstockler/omaha-order-manager-service-go/db"
	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

// armazena ordem no banco de dados
func storeOrder(o *models.Order) {
	db := db.MgoDb{}
	db.Open()
	err := db.Db.C("order").Insert(&o)
	db.Close()
	if err != nil {
		log.Printf("MONGO: %v", err.Error())
		o.SyncCode = 500
	}
}
