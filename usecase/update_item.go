package usecase

import (
	"encoding/json"
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/gorilla/mux"
)

// UpdateItem exported
func (u *UseCase) UpdateItem(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["order_id"]
	merchant := r.Header.Get("merchant_id")
	itemUUID := mux.Vars(r)["item_id"]

	order, err := cache.ShowOrder(merchant, uuid)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	newItem := models.Item{}
	json.NewDecoder(r.Body).Decode(&newItem)
	newItem.Build()

	for i := len(order.Items) - 1; i >= 0; i-- {
		if order.Items[i].UUID == itemUUID {
			newItem.UUID = order.Items[i].UUID
			order.Items = append(order.Items[:i], order.Items[i+1:]...)
			order.Items = append(order.Items, newItem)
		}
	}

	if err := u.DB.Update(*order); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, newItem)
}
