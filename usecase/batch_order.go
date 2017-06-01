package usecase

import (
	"encoding/json"
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

// BatchOrder exported
func (u *UseCase) BatchOrder(w http.ResponseWriter, r *http.Request) {

	merchant := r.Header.Get("merchant_id")
	logic := r.Header.Get("logic_number")
	orders := []models.Order{}

	// Populate the order data
	json.NewDecoder(r.Body).Decode(&orders)

	if len(orders) < 1 {
		respondWithJSON(w, http.StatusOK, orders)
		return
	}

	out := make(chan *models.Order)
	response := make(chan []models.Order)

	//response http
	go goResponse(w, out, len(orders), response)

	// processando ordens em paralelo
	splitOrders(orders, merchant, logic, out)

	syncResponse(w, response)
}

func syncResponse(w http.ResponseWriter,
	response chan []models.Order) {

	for {

		ret := <-response

		respondWithJSON(w, http.StatusOK, ret)

		return
	}

}

func goResponse(w http.ResponseWriter, out chan *models.Order,
	limit int, response chan []models.Order) {

	var ret []models.Order
	for {

		// Wait for the next job to come off the queue.
		receiveOrder := <-out

		ret = append(ret, *receiveOrder)

		if len(ret) == limit {
			break
		}
	}

	response <- ret
}

func splitOrders(orders []models.Order,
	merchant string,
	logic string,
	out chan *models.Order) {

	for i := 0; i < len(orders); i++ {
		go buildSyncOrder(&orders[i], merchant, logic, out)
	}
}
