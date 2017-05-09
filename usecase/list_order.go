package usecase

import (
	"log"
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
)

func listOrder(w http.ResponseWriter, r *http.Request) {

	merchantID := r.Header.Get("merchant_id")

	log.Printf("merchantID: %s", merchantID)

	o := rediscli.ListOrder(merchantID)

	respondWithJSON(w, http.StatusOK, o)
}
