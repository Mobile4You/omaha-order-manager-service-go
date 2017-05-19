package usecase

import (
	"strings"
	"time"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"gopkg.in/mgo.v2/bson"
)

// compila um item
func buildItem(i *models.Item) {
	if len(strings.TrimSpace(i.UUID.Hex())) == 0 {
		i.UUID = bson.NewObjectId()
		i.CreatedAt = time.Now()
	}
	i.UpdatedAt = time.Now()
}
