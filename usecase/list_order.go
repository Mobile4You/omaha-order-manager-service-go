package usecase

import (
	"net/http"
	"sort"
	"strings"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

// ListOrder exported
func (u *UseCase) ListOrder(w http.ResponseWriter, r *http.Request) {

	merchant := r.Header.Get("merchant_id")
	logic := r.Header.Get("logic_number")
	action := r.FormValue("action")

	if strings.TrimSpace(action) == "OPENED" {

		orders := cache.ListOrder(merchant)

		orders = openedOrder(orders, logic)

		sort.Sort(models.OrderAscending(orders))

		respondWithJSON(w, http.StatusOK, orders)

		return
	}

	respondWithJSON(w, http.StatusOK, closedOrder(u))
}

func closedOrder(u *UseCase) []models.JSONB{

	var dbOrders []models.OrderPg
	u.DB.Table("orders").Find(&dbOrders)

	orders := make([]models.JSONB, 0)

	for _, v := range dbOrders {
		orders = append(orders, v.Payload)
	}

	return orders
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
