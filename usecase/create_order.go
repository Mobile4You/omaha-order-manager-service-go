package usecase

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

// createOrder represent new order
func createOrder(w http.ResponseWriter, r *http.Request) {

	merchant := r.Header.Get("merchant_id")
	logic := r.Header.Get("logic_number")

	o := models.Order{}

	// Populate the order data
	json.NewDecoder(r.Body).Decode(&o)

	if len(o.Items) < 1 {
		respondWithError(w, http.StatusNotFound, errors.New("order without items").Error())
		return
	}

	buildOrder(&o, merchant, &logic)

	respondWithJSON(w, http.StatusOK, o)

}
