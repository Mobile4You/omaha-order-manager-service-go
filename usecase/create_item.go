package usecase

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/gorilla/mux"
)

// CreateItem exported
func (u *UseCase) CreateItem(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["order_id"]
	merchant := r.Header.Get("merchant_id")

	order, err := cache.ShowOrder(merchant, uuid)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	item := models.Item{}
	json.NewDecoder(r.Body).Decode(&item)

	item.Build()

	order.Items = append(order.Items, item)
	order.UpdatedAt = time.Now()

	if err := u.DB.Update(*order); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, item)
}
