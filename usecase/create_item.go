package usecase

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
	"github.com/gorilla/mux"
)

func createItem(w http.ResponseWriter, r *http.Request) {

	orderUUID := mux.Vars(r)["order_id"]
	merchantID := r.Header.Get("merchant_id")

	order, err := rediscli.FindOrder(merchantID, orderUUID)
	if err != nil || len(strings.TrimSpace(order.UUID.String())) == 0 {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	i := models.Item{}
	json.NewDecoder(r.Body).Decode(&i)

	// Add an UUID
	i.UUID = bson.NewObjectId()
	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()

	order.Items = append(order.Items, i)
	order.UpdatedAt = time.Now()

	err = rediscli.PutOrder(*order)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, i)
}
