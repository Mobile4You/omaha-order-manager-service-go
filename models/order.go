package models

import (
	"time"

	"github.com/go-redis/redis"

	"gopkg.in/mgo.v2/bson"
)

// DRAFT is an exported
const (
	DRAFT   OrderStatusType = "DRAFT"
	ENTERED OrderStatusType = "ENTERED"
	PAID    OrderStatusType = "PAID"
	CLOSED  OrderStatusType = "CLOSED"
)

type (

	// Terminal is an exported
	Terminal struct {
		UUID      bson.ObjectId `json:"id" bson:"_id"`
		Number    string        `json:"number" validate:"required"`
		CreatedAt time.Time     `json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt time.Time     `json:"updated_at,omitempty" bson:"updated_at"`
		Sub       *redis.PubSub
	}

	// OrderStatusType is an exported
	OrderStatusType string

	// Order represents the structure send lio
	Order struct {
		UUID        bson.ObjectId   `json:"id" bson:"_id"`
		Number      string          `json:"number,omitempty" bson:"number"`
		Reference   string          `json:"reference,omitempty" bson:"reference"`
		Notes       string          `json:"notes,omitempty" bson:"notes"`
		CreatedAt   time.Time       `json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt   time.Time       `json:"updated_at,omitempty" bson:"updated_at"`
		MerchantID  string          `json:"merchant_id" bson:"merchant_id" validate:"required"`
		LogicNumber string          `json:"logic_number" bson:"logic_number" validate:"required"`
		Status      OrderStatusType `json:"status" bson:"status" validate:"required"`
		Ref         string          `json:"ref"`
		Items       []Item          `json:"items,omitempty" bson:"items"`
		//Terminal     Terminal             `json:"terminal"`
		//Transactions []PaymentTransaction `json:"transactions,omitempty" bson:"xxxxxxxxxxxxx"`
		//Price      int    `json:"price,omitempty" bson:"price"`
	}

	// Item is an exported
	Item struct {
		UUID          bson.ObjectId `json:"id" bson:"_id"`
		CreatedAt     time.Time     `json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt     time.Time     `json:"updated_at,omitempty" bson:"updated_at"`
		Sku           string        `json:"sku,omitempty" bson:"sku"`
		SkuType       string        `json:"sku_type,omitempty" bson:"sku_type"`
		Name          string        `json:"name,omitempty" bson:"name"`
		Description   string        `json:"description,omitempty" bson:"description"`
		UnitPrice     string        `json:"unit_price,omitempty" bson:"unit_price"`
		Quantity      int           `json:"quantity,omitempty" bson:"quantity"`
		UnitOfMeasure string        `json:"unit_of_measure,omitempty" bson:"unit_of_measure"`
		Details       string        `json:"details,omitempty" bson:"details"`
	}
)
