package usecase

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

func TestListOrder(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/v3/orders?action=OPENED", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListOrder)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestFilterOrder(t *testing.T) {
	orders := []models.Order{}
	orders = append(orders, models.Order{})

	ret := filterOrder(orders, func(arg2 models.Order) bool {
		if arg2.Status == models.CLOSED {
			return false
		}
		return true
	})

	if len(ret) != 1 {
		t.Errorf("Expected size 1, but it was %d instead.", len(ret))
	}
}

func TestFilterOpenOrder(t *testing.T) {

	orders := []models.Order{}
	orders = append(orders, models.Order{LogicNumber: "133"})
	orders = append(orders, models.Order{LogicNumber: "122"})
	orders = append(orders, models.Order{LogicNumber: "133", Status: models.DRAFT})
	orders = append(orders, models.Order{LogicNumber: "133", Status: models.CLOSED})
	orders = append(orders, models.Order{LogicNumber: "122"})

	ret := openedOrder(orders, "133")

	if len(ret) != 2 {
		t.Errorf("Expected size 2, but it was %d instead.", len(ret))
	}
}
