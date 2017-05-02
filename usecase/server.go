package usecase

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	log.Println("iniciando http server ...")
	router := mux.NewRouter().StrictSlash(true)
	apiV3(router)
	http.ListenAndServe(":8080", router)
}

func apiV3(router *mux.Router) {
	apiV3 := router.PathPrefix("/api/v3").Subrouter()

	apiV3.Handle("/order", ensureBaseOrder(http.HandlerFunc(ListOrder))).Methods("GET")
	apiV3.Handle("/order", ensureBaseOrder(http.HandlerFunc(CreateOrder))).Methods("POST")
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithString(w http.ResponseWriter, code int, payload string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(payload))
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
