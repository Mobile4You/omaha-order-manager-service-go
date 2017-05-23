package usecase

import (
	"testing"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

func TestEmptyItems(t *testing.T) {
	o := models.Order{
		Items: make([]models.Item, 0),
	}

	if err := buildOrder(&o, "", ""); err == nil {
		t.Error("Expected order without items")
	}

}

func TestNewOrder(t *testing.T) {
	o := models.Order{
		Items: make([]models.Item, 0),
	}
	o.Items = append(o.Items, models.Item{})

	buildOrder(&o, "", "")

	if len(o.UUID.Hex()) == 0 {
		t.Error("Expected new UUID")
	}

	if o.CreatedAt.IsZero() || o.UpdatedAt.IsZero() {
		t.Error("Value of data fields, can not be null")
	}

	if o.SyncCode != 201 {
		t.Error("Expected http status 201")
	}
}