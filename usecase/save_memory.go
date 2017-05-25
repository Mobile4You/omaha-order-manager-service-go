package usecase

import (
	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
)

func saveMemory(o *models.Order) {
	re := rediscli.ORedis{}
	err := re.PutOrder(*o)
	if err != nil {
		o.SyncCode = 400
	}
}
