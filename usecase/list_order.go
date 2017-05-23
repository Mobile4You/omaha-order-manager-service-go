package usecase

import (
	"net/http"
	"sort"
	"strings"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
)

// ListOrder exported
func ListOrder(w http.ResponseWriter, r *http.Request) {

	merchant := r.Header.Get("merchant_id")
	logic := r.Header.Get("logic_number")
	action := r.FormValue("action")

	if strings.TrimSpace(action) == "OPENED" {

		orders := rediscli.ListOrder(merchant)

		orders = openedOrder(orders, logic)

		sort.Sort(models.OrderAscending(orders))

		respondWithJSON(w, http.StatusOK, orders)

		return
	}

	//TODO: CLOSED ORDER
	respondWithJSON(w, http.StatusOK, action)
}

func openedOrder(orders []models.Order, number string) []models.Order {

	testOpen := func(o models.Order) bool {
		if o.Status == models.CLOSED {
			return false
		}

		if len(strings.TrimSpace(number)) == 0 {
			return true
		}
		return o.LogicNumber == number
	}

	return filterOrder(orders, testOpen)
}

func filterOrder(orders []models.Order, test func(models.Order) bool) (ret []models.Order) {
	for _, o := range orders {
		if test(o) {
			ret = append(ret, o)
		}
	}
	return
}
