package routes

import (
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/usecase"
	"github.com/gorilla/mux"
)

// Router exported
type Router struct {
	Use *usecase.UseCase
}

// ServerStart is exported
func (r *Router) ServerStart() {
	router := mux.NewRouter().StrictSlash(true)
	apiV3(r, router)
	http.ListenAndServe(":8080", router)
}

func apiV3(r *Router, router *mux.Router) {
	api := router.PathPrefix("/api/v3").Subrouter()
	apiOrder(r, api)
	apiItem(r, api)
	apiTransaction(r, api)
}

func apiOrder(r *Router, api *mux.Router) {
	api.Handle("/orders", ensureHeader(http.HandlerFunc(r.Use.CreateOrder))).Methods("POST")
	api.Handle("/orders", ensureHeader(http.HandlerFunc(r.Use.ListOrder))).Methods("GET")
	api.Handle("/orders/{order_id}", ensureHeader(http.HandlerFunc(r.Use.UpdateOrder))).Methods("PUT")
	api.Handle("/orders/{order_id}", ensureHeader(http.HandlerFunc(r.Use.ShowOrder))).Methods("GET")
	api.Handle("/orders/batch", ensureHeader(http.HandlerFunc(r.Use.BatchOrder))).Methods("POST")
}

func apiItem(r *Router, api *mux.Router) {
	api.Handle("/orders/{order_id}/items", ensureHeader(http.HandlerFunc(r.Use.CreateItem))).Methods("POST")
	api.Handle("/orders/{order_id}/items/{item_id}", ensureHeader(http.HandlerFunc(r.Use.UpdateItem))).Methods("PUT")
	api.Handle("/orders/{order_id}/items/{item_id}", ensureHeader(http.HandlerFunc(r.Use.ShowItem))).Methods("GET")
	api.Handle("/orders/{order_id}/items/{item_id}", ensureHeader(http.HandlerFunc(r.Use.DeleteItem))).Methods("DELETE")
}

func apiTransaction(r *Router, api *mux.Router) {
	api.Handle("/orders/{order_id}/transactions", ensureHeader(http.HandlerFunc(r.Use.CreateTransaction))).Methods("POST")
	api.Handle("/orders/{order_id}/transactions", ensureHeader(http.HandlerFunc(r.Use.ListTransaction))).Methods("GET")
}
