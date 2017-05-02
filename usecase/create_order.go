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

// CreateOrder represent new order
func CreateOrder(w http.ResponseWriter, r *http.Request) {

	o := models.Order{}

	// Populate the order data
	json.NewDecoder(r.Body).Decode(&o)

	// Add an UUID
	o.UUID = bson.NewObjectId()
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()

	if o.Status != models.CLOSED {
		memOrder(w, r, o)
		return
	}

	closeOrder(w, r, o)
}

// Closed order for mongodb
func closeOrder(w http.ResponseWriter, r *http.Request, o models.Order) {

	db := db.MgoDb{}

	db.Open()

	err := db.Db.C("order").Insert(o)

	db.Close()

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, o)
}

// Transactional order (status DRAFT, PAID, ENTERED)
func memOrder(w http.ResponseWriter, r *http.Request, o models.Order) {

	err := rediscli.PutOrder(o)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, o)
}
