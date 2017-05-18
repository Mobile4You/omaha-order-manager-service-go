package usecase

import (
	"log"
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
	"github.com/gorilla/mux"
)

func shareOrder(w http.ResponseWriter, r *http.Request) {

	merchantID := r.Header.Get("merchant_id")
	orderUUID := mux.Vars(r)["order_id"]

	ch, err := rediscli.CreateChannel(merchantID, orderUUID)

	log.Printf("channel: %v , error: %v", ch, err)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, ch)
}
