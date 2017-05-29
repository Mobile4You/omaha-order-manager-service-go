package usecase

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ShowOrder represent one order
func (u *UseCase) ShowOrder(w http.ResponseWriter, r *http.Request) {

	merchant := r.Header.Get("merchant_id")
	uuid := mux.Vars(r)["order_id"]

	order, err := cache.ShowOrder(merchant, uuid)
	if err == nil && order != nil {
		respondWithJSON(w, http.StatusOK, order)
		return
	}

	//bsuca no banco de dados
	respondWithJSON(w, http.StatusOK, order)
}
