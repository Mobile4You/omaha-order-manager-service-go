package usecase

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// RemoteAPI exported
type RemoteAPI interface {
	//feito
	CreateOrder(w http.ResponseWriter, r *http.Request)
	//feito
	ListOrder(w http.ResponseWriter, r *http.Request)
	//feito
	UpdateOrder(w http.ResponseWriter, r *http.Request)
	//feito
	ShowOrder(w http.ResponseWriter, r *http.Request)
	DeleteOrder(w http.ResponseWriter, r *http.Request)
	//feito
	CreateItem(w http.ResponseWriter, r *http.Request)
	//feito
	UpdateItem(w http.ResponseWriter, r *http.Request)
	//feito
	ShowItem(w http.ResponseWriter, r *http.Request)
	//feito
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
	apiSync(api)
}

func apiOrder(api *mux.Router) {
	api.Handle("/orders", ensureBaseOrder(http.HandlerFunc(use.CreateOrder))).Methods("POST")
	api.Handle("/orders", ensureBaseOrder(http.HandlerFunc(use.ListOrder))).Methods("GET")
	api.Handle("/orders/{order_id}", ensureBaseOrder(http.HandlerFunc(use.UpdateOrder))).Methods("PUT")
	api.Handle("/orders/{order_id}", ensureBaseOrder(http.HandlerFunc(use.ShowOrder))).Methods("GET")
}

func apiItem(api *mux.Router) {
	api.Handle("/orders/{order_id}/items", ensureBaseOrder(http.HandlerFunc(use.CreateItem))).Methods("POST")
	api.Handle("/orders/{order_id}/items/{item_id}", ensureBaseOrder(http.HandlerFunc(use.UpdateItem))).Methods("PUT")
	api.Handle("/orders/{order_id}/items/{item_id}", ensureBaseOrder(http.HandlerFunc(use.ShowItem))).Methods("GET")
	api.Handle("/orders/{order_id}/items/{item_id}", ensureBaseOrder(http.HandlerFunc(use.DeleteItem))).Methods("DELETE")
}

func apiTransaction(api *mux.Router) {

}

func apiSync(api *mux.Router) {
	// api.Handle("/sync", ensureBaseOrder(http.HandlerFunc(syncOrder))).Methods("POST")
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
