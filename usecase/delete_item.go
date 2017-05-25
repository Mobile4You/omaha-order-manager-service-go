package usecase

import (
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
	"github.com/gorilla/mux"
)

func deleteItem(w http.ResponseWriter, r *http.Request) {

	orderUUID := mux.Vars(r)["order_id"]
	merchant := r.Header.Get("merchant_id")
	itemUUID := mux.Vars(r)["item_id"]

	order, err := rediscli.FindOrder(merchant, orderUUID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	for i := len(order.Items) - 1; i >= 0; i-- {
		if order.Items[i].UUID.Hex() == itemUUID {
			order.Items = append(order.Items[:i], order.Items[i+1:]...)
		}
	}

	re := rediscli.ORedis{}
	if err = re.PutOrder(*order); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithCode(w, http.StatusOK)
}
