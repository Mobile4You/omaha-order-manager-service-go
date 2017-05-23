package usecase

import (
	"encoding/json"
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

// CreateOrder represent new order
func CreateOrder(w http.ResponseWriter, r *http.Request) {

	merchant := r.Header.Get("merchant_id")
	logic := r.Header.Get("logic_number")

	o := models.Order{}

	// Populate the order data
	json.NewDecoder(r.Body).Decode(&o)

	if err := buildOrder(&o, merchant, logic); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, o)
}
