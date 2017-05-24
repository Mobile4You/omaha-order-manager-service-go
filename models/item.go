package models

import (
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (

	// Item is an exported
	Item struct {
		Build         interface{}   `json:"_,omitempty"`
		UUID          bson.ObjectId `json:"id" bson:"_id"`
		CreatedAt     time.Time     `json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt     time.Time     `json:"updated_at,omitempty" bson:"updated_at"`
		Sku           string        `json:"sku,omitempty" bson:"sku"`
		SkuType       string        `json:"sku_type,omitempty" bson:"sku_type"`
		Name          string        `json:"name,omitempty" bson:"name"`
		Description   string        `json:"description,omitempty" bson:"description"`
		UnitPrice     int           `json:"unit_price,omitempty" bson:"unit_price"`
		Quantity      int           `json:"quantity,omitempty" bson:"quantity"`
		UnitOfMeasure string        `json:"unit_of_measure,omitempty" bson:"unit_of_measure"`
		Details       string        `json:"details,omitempty" bson:"details"`
		Ref           string        `json:"ref"`
	}
)

func (i *Item) Run(args ...interface{}) error {
	if len(strings.TrimSpace(i.UUID.Hex())) == 0 {
		i.UUID = bson.NewObjectId()
		i.CreatedAt = time.Now()
	}
	i.UpdatedAt = time.Now()
	return nil
}
