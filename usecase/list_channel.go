package usecase

import (
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
)

func listChannel(w http.ResponseWriter, r *http.Request) {

	merchantID := r.Header.Get("merchant_id")

	ch := rediscli.ListChannel(merchantID)

	respondWithJSON(w, http.StatusOK, ch)
}
