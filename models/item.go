package models

import (
	"strings"
	"time"

	"github.com/satori/go.uuid"
)

type (

	// Item is an exported
	Item struct {
		BuildModel    interface{} `json:"_,omitempty"`
		UUID          string      `json:"id"`
		CreatedAt     time.Time   `json:"created_at,omitempty"`
		UpdatedAt     time.Time   `json:"updated_at,omitempty"`
		Sku           string      `json:"sku,omitempty"`
		SkuType       string      `json:"sku_type,omitempty"`
		Name          string      `json:"name,omitempty"`
		Description   string      `json:"description,omitempty"`
		UnitPrice     int         `json:"unit_price,omitempty"`
		Quantity      int         `json:"quantity,omitempty"`
		UnitOfMeasure string      `json:"unit_of_measure,omitempty"`
		Details       string      `json:"details,omitempty"`
		Ref           string      `json:"ref"`
	}
)

// Build exported Item
func (i *Item) Build(args ...interface{}) error {
	if len(strings.TrimSpace(i.UUID)) == 0 {
		i.UUID = uuid.NewV4().String()
		i.CreatedAt = time.Now()
	}
	i.UpdatedAt = time.Now()
	return nil
}
