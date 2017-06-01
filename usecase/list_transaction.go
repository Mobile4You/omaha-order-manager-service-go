package usecase

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ListTransaction exported
func (u *UseCase) ListTransaction(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["order_id"]
	merchant := r.Header.Get("merchant_id")

	order, err := cache.ShowOrder(merchant, uuid)
	if err == nil && order != nil {
		respondWithJSON(w, http.StatusOK, order.Transactions)
		return
	}

	respondWithCode(w, http.StatusBadGateway)
}
