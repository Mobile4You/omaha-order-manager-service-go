package usecase

import (
	"encoding/json"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/arthurstockler/omaha-order-manager-service-go/db"
	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
)

// createOrder represent new order
func createOrder(w http.ResponseWriter, r *http.Request) {

	merchantID := r.Header.Get("merchant_id")
	logicNumber := r.Header.Get("logic_number")

	o := models.Order{}

	// Populate the order data
	json.NewDecoder(r.Body).Decode(&o)

	// Add an UUID
	o.UUID = bson.NewObjectId()
	o.MerchantID = merchantID
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
	o.LogicNumber = logicNumber

	if len(o.Items) > 0 {
		o.Items[0].UUID = bson.NewObjectId()
	}

	if o.Status != models.CLOSED {
		memOrder(w, r, o)
		return
	}

	saveOrder(w, r, o)
}

// Closed order for mongodb
func saveOrder(w http.ResponseWriter, r *http.Request, o models.Order) {

	db := db.MgoDb{}

	db.Open()

	err := db.Db.C("order").Insert(o)

	db.Close()

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, o)
}

// Transactional order (status DRAFT, PAID, ENTERED)
func memOrder(w http.ResponseWriter, r *http.Request, o models.Order) {

	err := rediscli.PutOrder(o)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, o)
}
