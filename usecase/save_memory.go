package usecase

import (
	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
)

func saveMemory(o *models.Order) {
	err := rediscli.PutOrder(*o)
	if err != nil {
		o.SyncCode = 400
		return
	}
	o.SyncCode = 200
}
