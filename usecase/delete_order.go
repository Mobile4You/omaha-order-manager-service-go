package usecase

import (
	"errors"
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
	"github.com/gorilla/mux"
)

// DeleteOrder only order in DRAFT
func DeleteOrder(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["order_id"]
	merchant := r.Header.Get("merchant_id")

	order, err := rediscli.FindOrder(merchant, uuid)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	if order.Status != models.DRAFT {
		respondWithError(w, http.StatusBadRequest, errors.New("Only DRAFT orders can be deleted").Error())
		return
	}

	if err = rediscli.DelOrder(*order); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithCode(w, http.StatusOK)
}
