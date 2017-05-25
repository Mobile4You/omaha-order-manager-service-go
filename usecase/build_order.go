package usecase

import "github.com/arthurstockler/omaha-order-manager-service-go/models"

func buildSyncOrder(o *models.Order, merchant string, logic string,
	out chan *models.Order) {

	//ordem sem nenhum item não deveria ser processada
	if len(o.Items) < 1 {
		o.SyncCode = 400
		out <- o
		return
	}

	o.Build(merchant, logic)

	out <- o
}
