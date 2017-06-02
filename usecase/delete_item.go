package usecase

import (
	"net/http"

	"github.com/gorilla/mux"
)

// DeleteItem exported
func (u *UseCase) DeleteItem(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["order_id"]
	merchant := r.Header.Get("merchant_id")
	itemUUID := mux.Vars(r)["item_id"]

	order, err := cache.ShowOrder(merchant, uuid)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	for i := len(order.Items) - 1; i >= 0; i-- {
		if order.Items[i].UUID == itemUUID {
			order.Items = append(order.Items[:i], order.Items[i+1:]...)
		}
	}

	if err := u.DB.Update(*order); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithCode(w, http.StatusNoContent)
}
