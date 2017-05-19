package usecase

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
	"github.com/gorilla/mux"
)

func updateOrder(w http.ResponseWriter, r *http.Request) {

	merchant := r.Header.Get("merchant_id")
	orderUUID := mux.Vars(r)["order_id"]
	newOrder := models.Order{}

	json.NewDecoder(r.Body).Decode(&newOrder)

	if len(newOrder.Items) < 1 {
		respondWithError(w, http.StatusNotFound, errors.New("order without items").Error())
		return
	}

	o, err := rediscli.FindOrder(merchant, orderUUID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	newOrder.UUID = o.UUID
	newOrder.CreatedAt = o.CreatedAt
	newOrder.LogicNumber = o.LogicNumber

	buildOrder(&newOrder, merchant, nil)

	respondWithJSON(w, http.StatusCreated, newOrder)
}
