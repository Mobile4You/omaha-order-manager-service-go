package usecase

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
}

// ServerStart is exported
func ServerStart() {
	log.Println("iniciando http server ...")
	router := mux.NewRouter().StrictSlash(true)
	apiV3(router)
	http.ListenAndServe(":8080", router)
}

func apiV3(router *mux.Router) {
	api := router.PathPrefix("/api/v3").Subrouter()
	apiOrder(api)
	apiItem(api)
	apiTransaction(api)
}

func apiOrder(api *mux.Router) {
	api.Handle("/orders", ensureHeader(http.HandlerFunc(use.CreateOrder))).Methods("POST")
	api.Handle("/orders", ensureHeader(http.HandlerFunc(use.ListOrder))).Methods("GET")
	api.Handle("/orders/{order_id}", ensureHeader(http.HandlerFunc(use.UpdateOrder))).Methods("PUT")
	api.Handle("/orders/{order_id}", ensureHeader(http.HandlerFunc(use.ShowOrder))).Methods("GET")
	api.Handle("/orders/batch", ensureHeader(http.HandlerFunc(use.BatchOrder))).Methods("POST")
}

func apiItem(api *mux.Router) {
	api.Handle("/orders/{order_id}/items", ensureHeader(http.HandlerFunc(use.CreateItem))).Methods("POST")
	api.Handle("/orders/{order_id}/items/{item_id}", ensureHeader(http.HandlerFunc(use.UpdateItem))).Methods("PUT")
	api.Handle("/orders/{order_id}/items/{item_id}", ensureHeader(http.HandlerFunc(use.ShowItem))).Methods("GET")
	api.Handle("/orders/{order_id}/items/{item_id}", ensureHeader(http.HandlerFunc(use.DeleteItem))).Methods("DELETE")
}

func apiTransaction(api *mux.Router) {
	api.Handle("/orders/{order_id}/transactions", ensureHeader(http.HandlerFunc(use.CreateTransaction))).Methods("POST")
	api.Handle("/orders/{order_id}/transactions", ensureHeader(http.HandlerFunc(use.ListTransaction))).Methods("GET")
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
	w.Write(response)
}
