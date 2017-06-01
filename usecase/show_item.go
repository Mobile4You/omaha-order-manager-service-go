package usecase

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ShowItem exported
func (u *UseCase) ShowItem(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["order_id"]
	merchant := r.Header.Get("merchant_id")
	itemUUID := mux.Vars(r)["item_id"]

	order, err := cache.ShowOrder(merchant, uuid)

	if err == nil && order != nil {
		for i := len(order.Items) - 1; i >= 0; i-- {
			if order.Items[i].UUID == itemUUID {
				respondWithJSON(w, http.StatusOK, order.Items[i])
				return
			}
		}
		respondWithCode(w, http.StatusNoContent)
		return
	}

	//TODO: BUSCAR NO BANCO DE DADOS
	respondWithCode(w, http.StatusOK)
}
