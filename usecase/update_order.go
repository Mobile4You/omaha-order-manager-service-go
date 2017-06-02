package usecase

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/gorilla/mux"
)

// UpdateOrder exported
func (u *UseCase) UpdateOrder(w http.ResponseWriter, r *http.Request) {

	merchant := r.Header.Get("merchant_id")
	uuid := mux.Vars(r)["order_id"]
	newOrder := models.Order{}

	json.NewDecoder(r.Body).Decode(&newOrder)

	if len(newOrder.Items) < 1 {
		respondWithError(w, http.StatusNotFound, errors.New("order without items").Error())
		return
	}

	o, err := cache.ShowOrder(merchant, uuid)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	newOrder.UUID = o.UUID
	newOrder.Status = o.Status
	newOrder.CreatedAt = o.CreatedAt
	newOrder.LogicNumber = o.LogicNumber

	if err := newOrder.Build(merchant, o.LogicNumber); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := u.DB.Update(newOrder); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, newOrder)
}
