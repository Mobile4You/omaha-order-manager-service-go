package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

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
	// BuildModel is an exported
	BuildModel interface {
		Build(args ...interface{}) error
	}

	// OrderStatusType is an exported
	OrderStatusType string

	// Order represents the structure send lio
	Order struct {
		BuildModel   interface{}     `json:"_,omitempty"`
		UUID         bson.ObjectId   `json:"id" bson:"_id"`
		Number       string          `json:"number,omitempty" bson:"number"`
		Reference    string          `json:"reference,omitempty" bson:"reference"`
		Notes        string          `json:"notes,omitempty" bson:"notes"`
		CreatedAt    time.Time       `json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt    time.Time       `json:"updated_at,omitempty" bson:"updated_at"`
		MerchantID   string          `json:"merchant_id" bson:"merchant_id" valid:"required"`
		LogicNumber  string          `json:"logic_number" bson:"logic_number" valid:"required"`
		Status       OrderStatusType `json:"status" bson:"status" valid:"required"`
		Ref          string          `json:"ref"`
		SyncCode     int             `json:"sync_code"`
		Items        []Item          `json:"items,omitempty" bson:"items"`
		Transactions []Transaction   `json:"transactions,omitempty" bson:"transactions"`
	}

	// OrderAscending represents the structure Ordered
	OrderAscending []Order
)

func (v OrderAscending) Len() int           { return len(v) }
func (v OrderAscending) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v OrderAscending) Less(i, j int) bool { return v[j].CreatedAt.After(v[i].CreatedAt) }

// Build exported Order
func (o *Order) Build(args ...interface{}) error {
	if len(o.Items) < 1 {
		return errors.New("order without items")
	}
	if len(o.Status) == 0 {
		o.Status = DRAFT
	}
	o.SyncCode = 200
	if len(strings.TrimSpace(o.UUID.Hex())) == 0 {
		o.UUID = bson.NewObjectId()
		o.CreatedAt = time.Now()
		o.SyncCode = 201
	}
	o.UpdatedAt = time.Now()
	o.MerchantID = reflect.ValueOf(args).Index(0).Interface().(string)
	o.LogicNumber = reflect.ValueOf(args).Index(1).Interface().(string)

	for i := 0; i < len(o.Items); i++ {
		o.Items[i].Build()
	}

	return nil
}
