package usecase

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Init is exported
func Init() {
	log.Println("iniciando http server ...")
	router := mux.NewRouter().StrictSlash(true)
	apiV3(router)
	http.ListenAndServe(":8080", router)
}

func apiV3(router *mux.Router) {
	api := router.PathPrefix("/api/v3").Subrouter()
	apiOrder(api)
	apiItem(api)
	apiChannel(api)
}

func apiOrder(api *mux.Router) {
	api.Handle("/orders", ensureBaseOrder(http.HandlerFunc(listOrder))).Methods("GET")
	api.Handle("/orders", ensureBaseOrder(http.HandlerFunc(createOrder))).Methods("POST")
	api.Handle("/orders/{order_id}", ensureBaseOrder(http.HandlerFunc(updateOrder))).Methods("PUT")
	api.Handle("/orders/{order_id}/share", ensureBaseOrder(http.HandlerFunc(shareOrder))).Methods("PUT")
}

func apiItem(api *mux.Router) {
	api.Handle("/orders/{order_id}/items", ensureBaseOrder(http.HandlerFunc(createItem))).Methods("POST")
	api.Handle("/orders/{order_id}/items/{item_id}", ensureBaseOrder(http.HandlerFunc(deleteItem))).Methods("DELETE")
	api.Handle("/orders/{order_id}/items/{item_id}", ensureBaseOrder(http.HandlerFunc(updateItem))).Methods("PUT")
}

func apiChannel(api *mux.Router) {
	api.Handle("/channel", ensureBaseOrder(http.HandlerFunc(listChannel))).Methods("GET")
	api.Handle("/channel/{channel_id}/subscribe", ensureBaseOrder(http.HandlerFunc(subscribeChannel))).Methods("GET")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithString(w http.ResponseWriter, code int, payload string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(payload))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
