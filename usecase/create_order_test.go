package usecase

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

func TestCreateOrder(t *testing.T) {
	o := models.Order{Items: make([]models.Item, 0)}
	o.Items = append(o.Items, models.Item{})

	body, _ := json.Marshal(o)
	req, err := http.NewRequest("POST", "/api/v3/orders", bytes.NewReader(body))
	req.Header.Add("merchant_id", "123")
	req.Header.Add("logic_number", "777777-9")

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateOrder)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
