package usecase

import (
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
	"github.com/gorilla/mux"
)

func deleteOrder(w http.ResponseWriter, r *http.Request) {

	orderUUID := mux.Vars(r)["order_id"]
	merchantID := r.Header.Get("merchant_id")

	order, err := rediscli.FindOrder(merchantID, orderUUID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	if err = rediscli.DelOrder(*order); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithCode(w, http.StatusOK)
}
