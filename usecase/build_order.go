package usecase

import (
	"errors"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

func buildOrder(o *models.Order, merchant string, logic string) error {
	if len(o.Items) < 1 {
		return errors.New("order without items")
	}

	o.SyncCode = 200
	if len(strings.TrimSpace(o.UUID.Hex())) == 0 {
		o.UUID = bson.NewObjectId()
		o.CreatedAt = time.Now()
		o.SyncCode = 201
	}
	o.LogicNumber = logic
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

	return nil
}

func buildSyncOrder(o *models.Order, merchant string, logic string,
	out chan *models.Order) {

	//ordem sem nenhum item não deveria ser processada
	if len(o.Items) < 1 {
		o.SyncCode = 400
		out <- o
		return
	}

	buildOrder(o, merchant, logic)

	out <- o
}
