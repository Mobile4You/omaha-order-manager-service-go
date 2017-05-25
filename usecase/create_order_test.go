package usecase

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

func testAddRequest(t *testing.T, req *http.Request, test func(int)) {
	rr := httptest.NewRecorder()
	use := &UseCase{}
	req.Header.Add("merchant_id", "123")
	req.Header.Add("logic_number", "777777-9")
	handler := http.HandlerFunc(use.CreateOrder)
	handler.ServeHTTP(rr, req)
}

func TestCreateOrderWithoutItems(t *testing.T) {
	o := models.Order{Items: make([]models.Item, 0)}
	body, _ := json.Marshal(o)
	req, err := http.NewRequest("POST", "/api/v3/orders", bytes.NewReader(body))
	if err != nil {
		t.Error(err)
	}
	testAddRequest(t, req, func(status int) {
		if status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})
}

func TestCreateOrder(t *testing.T) {
	o := models.Order{Items: make([]models.Item, 0)}
	o.Items = append(o.Items, models.Item{})
	body, _ := json.Marshal(o)
	req, err := http.NewRequest("POST", "/api/v3/orders", bytes.NewReader(body))
	if err != nil {
		t.Error(err)
	}
	testAddRequest(t, req, func(status int) {
		if status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

}
