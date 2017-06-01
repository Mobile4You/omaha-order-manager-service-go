package usecase

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
)

// CreateTransaction exported
func (u *UseCase) CreateTransaction(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["order_id"]
	merchant := r.Header.Get("merchant_id")

	transaction := models.Transaction{}
	json.NewDecoder(r.Body).Decode(&transaction)
	transaction.Build()

	_, err := govalidator.ValidateStruct(transaction)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	order, err := cache.ShowOrder(merchant, uuid)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	order.Transactions = append(order.Transactions, transaction)
	order.UpdatedAt = time.Now()

	if err := u.SaveOrder(*order); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, transaction)
}
