package usecase

import (
	"net/http"
	"net/http/httptest"
	"testing"
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
