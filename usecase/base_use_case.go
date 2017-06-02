package usecase

import (
	"encoding/json"
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/caching"
	"github.com/jinzhu/gorm"
)

// RemoteAPI exported
type RemoteAPI interface {
	BatchOrder(w http.ResponseWriter, r *http.Request)
	CreateOrder(w http.ResponseWriter, r *http.Request)
	ListOrder(w http.ResponseWriter, r *http.Request)
	UpdateOrder(w http.ResponseWriter, r *http.Request)
	ShowOrder(w http.ResponseWriter, r *http.Request)
	DeleteOrder(w http.ResponseWriter, r *http.Request)
	CreateItem(w http.ResponseWriter, r *http.Request)
	UpdateItem(w http.ResponseWriter, r *http.Request)
	ShowItem(w http.ResponseWriter, r *http.Request)
	DeleteItem(w http.ResponseWriter, r *http.Request)
	CreateTransaction(w http.ResponseWriter, r *http.Request)
	ListTransaction(w http.ResponseWriter, r *http.Request)
	OperationOrder(w http.ResponseWriter, r *http.Request)
}

var (
	cache caching.RedisCache
)

// Store exported
type Store struct {
	Conn *gorm.DB
	Persistence
}

// UseCase exported
type UseCase struct {
	DB *Store
	RemoteAPI
}

// Init exported
func (u *UseCase) init() {
	cache = caching.RedisCache{}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithString(w http.ResponseWriter, code int, payload string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(payload))
}

func respondWithCode(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if string(response) != "null" {
		w.Write(response)
	}
}
