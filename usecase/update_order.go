package usecase

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
	"github.com/gorilla/mux"
)

func updateOrder(w http.ResponseWriter, r *http.Request) {

	merchantID := r.Header.Get("merchant_id")
	orderUUID := mux.Vars(r)["order_id"]
	newOrder := models.Order{}

	// Populate the order data
	json.NewDecoder(r.Body).Decode(&newOrder)

	_, err := rediscli.FindOrder(merchantID, orderUUID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
	}

	newOrder.UpdatedAt = time.Now()

	if err := closeWithoutItems(newOrder); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	//total de transactions pagas nao pode ser menor que total de somatorio dos valores dos itens
	if newOrder.Status == models.CLOSED {
		calculatePayment(newOrder)
		rediscli.DelOrder(newOrder)
		saveOrder(w, r, newOrder)
		return
	}

	rediscli.PutOrder(newOrder)

	respondWithJSON(w, http.StatusCreated, newOrder)
}

func calculatePayment(o models.Order) error {
	return nil
}

func closeWithoutItems(o models.Order) error {
	if len(o.Items) == 0 && o.Status == models.CLOSED {
		return errors.New("Not able to close orders without items")
	}
	return nil
}
