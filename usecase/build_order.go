package usecase

import (
	"strings"
	"time"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"gopkg.in/mgo.v2/bson"
)

func buildOrder(o *models.Order, merchant string, logic *string) {

	if len(strings.TrimSpace(o.UUID.Hex())) == 0 {
		o.UUID = bson.NewObjectId()
		o.CreatedAt = time.Now()
		o.SyncCode = 201
	} else {
		o.SyncCode = 200
	}
	if logic != nil {
		o.LogicNumber = *logic
	}

	o.UpdatedAt = time.Now()
	o.MerchantID = merchant

	for i := 0; i < len(o.Items); i++ {
		buildItem(&o.Items[i])
	}

	if o.Status == models.CLOSED {
		storeOrder(o)
	} else {
		saveMemory(o)
	}

}

func buildSyncOrder(o *models.Order, merchant string, logic string,
	out chan *models.Order) {

	//ordem sem nenhum item não deveria ser processada
	if len(o.Items) < 1 {
		o.SyncCode = 400
		out <- o
		return
	}

	buildOrder(o, merchant, &logic)

	out <- o
}
