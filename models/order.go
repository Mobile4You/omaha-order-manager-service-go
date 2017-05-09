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
		//Terminal     Terminal             `json:"terminal"`
		//Items        []Item               `json:"items,omitempty" bson:"xxxxxxxxxxxxx"`
		//Transactions []PaymentTransaction `json:"transactions,omitempty" bson:"xxxxxxxxxxxxx"`
		//Price      int    `json:"price,omitempty" bson:"price"`

	}
)
