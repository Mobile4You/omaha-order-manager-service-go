package models

import (
	"strings"
	"time"

	"github.com/satori/go.uuid"
)

type (

	// Transaction is an exported
	Transaction struct {
		BuildModel        interface{}    `json:"_,omitempty"`
		UUID              string         `json:"id"`
		CreatedAt         time.Time      `json:"created_at,omitempty"`
		UpdatedAt         time.Time      `json:"-"`
		ExternalID        string         `json:"external_id,omitempty"`
		TransactionType   string         `json:"transaction_type,omitempty"`
		Status            string         `json:"status,omitempty" valid:"required"`
		Description       string         `json:"description,omitempty"`
		TerminalNumber    string         `json:"terminal_number,omitempty" valid:"required"`
		Number            string         `json:"number,omitempty" valid:"required"`
		AuthorizationCode string         `json:"authorization_code,omitempty" valid:"required"`
		Amount            int            `json:"amount,omitempty" valid:"required"`
		PaymentFields     string         `json:"payment_fields,omitempty" valid:"required"`
		PaymentProduct    PaymentProduct `json:"payment_product,omitempty"`
		Card              Card           `json:"card,omitempty"`
	}

	//PaymentProduct exported
	PaymentProduct struct {
		Number         int            `json:"number,omitempty" valid:"required"`
		Name           string         `json:"name,omitempty" valid:"required"`
		PaymentService PaymentService `json:"sub,omitempty"`
	}

	//PaymentService exported
	PaymentService struct {
		Number int    `json:"number,omitempty" valid:"required"`
		Name   string `json:"name,omitempty" valid:"required"`
	}

	//Card exported
	Card struct {
		Brand string `json:"brand,omitempty" valid:"required"`
		Mask  string `json:"mask,omitempty" valid:"required"`
	}
)

// Build exported Transaction
func (t *Transaction) Build(args ...interface{}) error {
	if len(strings.TrimSpace(t.UUID)) == 0 {
		t.UUID = uuid.NewV4().String()
		t.CreatedAt = time.Now()
	}
	t.UpdatedAt = time.Now()
	return nil
}
