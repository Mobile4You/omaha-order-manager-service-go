package usecase

import (
	"net/http"
)

// OperationOrder exported
func (u *UseCase) OperationOrder(w http.ResponseWriter, r *http.Request) {

	// uuid := mux.Vars(r)["order_id"]
	//status := mux.Vars(r)["status"]
	// merchant := r.Header.Get("merchant_id")

	// a := models.OrderStatus(status)
	//
	// fmt.Printf("format: %v", a)

	respondWithCode(w, http.StatusOK)
}
