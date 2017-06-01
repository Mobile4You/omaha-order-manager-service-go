package models

import (
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (

	// Transaction is an exported
	Transaction struct {
		BuildModel        interface{}    `json:"_,omitempty"`
		UUID              bson.ObjectId  `json:"id" bson:"_id"`
		CreatedAt         time.Time      `json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt         time.Time      `json:"-" bson:"updated_at"`
		ExternalID        string         `json:"external_id,omitempty" bson:"external_id"`
		TransactionType   string         `json:"transaction_type,omitempty" bson:"transaction_type"`
		Status            string         `json:"status,omitempty" bson:"status" valid:"required"`
		Description       string         `json:"description,omitempty" bson:"description"`
		TerminalNumber    string         `json:"terminal_number,omitempty" bson:"terminal_number" valid:"required"`
		Number            string         `json:"number,omitempty" bson:"number" valid:"required"`
		AuthorizationCode string         `json:"authorization_code,omitempty" bson:"authorization_code" valid:"required"`
		Amount            int            `json:"amount,omitempty" bson:"amount" valid:"required"`
		PaymentFields     string         `json:"payment_fields,omitempty" bson:"payment_fields" valid:"required"`
		PaymentProduct    PaymentProduct `json:"payment_product,omitempty" bson:"payment_product"`
		Card              Card           `json:"card,omitempty" bson:"card"`
	}

	//PaymentProduct exported
	PaymentProduct struct {
		Number         int            `json:"number,omitempty" bson:"number" valid:"required"`
		Name           string         `json:"name,omitempty" bson:"name" valid:"required"`
		PaymentService PaymentService `json:"sub,omitempty" bson:"sub"`
	}

	//PaymentService exported
	PaymentService struct {
		Number int    `json:"number,omitempty" bson:"number" valid:"required"`
		Name   string `json:"name,omitempty" bson:"name" valid:"required"`
	}

	//Card exported
	Card struct {
		Brand string `json:"brand,omitempty" bson:"brand" valid:"required"`
		Mask  string `json:"mask,omitempty" bson:"mask" valid:"required"`
	}
)

// Build exported Transaction
func (t *Transaction) Build(args ...interface{}) error {
	if len(strings.TrimSpace(t.UUID.Hex())) == 0 {
		t.UUID = bson.NewObjectId()
		t.CreatedAt = time.Now()
	}
	t.UpdatedAt = time.Now()
	return nil
}
