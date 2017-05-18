package usecase

import (
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
	"github.com/gorilla/mux"
)

func updateItem(w http.ResponseWriter, r *http.Request) {

	orderUUID := mux.Vars(r)["order_id"]
	merchantID := r.Header.Get("merchant_id")
	itemUUID := mux.Vars(r)["item_id"]

	order, err := rediscli.FindOrder(merchantID, orderUUID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	for i := len(order.Items) - 1; i >= 0; i-- {
		if order.Items[i].UUID.Hex() == itemUUID {
			order.Items = append(order.Items[:i], order.Items[i+1:]...)
		}
	}

}
