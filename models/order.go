package models

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/satori/go.uuid"
)

// JSONB is an exported
type JSONB []byte

// OrderStatus is an exported
type OrderStatus int

// DRAFT is an exported
const (
	Draft OrderStatus = iota
	Entered
	Paid
	Closed
	ReEntered
)

var (

	// OrderStatusToValue is an exported
	OrderStatusToValue = map[string]OrderStatus{
		"DRAFT":      Draft,
		"ENTERED":    Entered,
		"PAID":       Paid,
		"CLOSED":     Closed,
		"RE_ENTERED": ReEntered,
	}

	// OrderStatusToName is an exported
	OrderStatusToName = map[OrderStatus]string{
		Draft:     "DRAFT",
		Entered:   "ENTERED",
		Paid:      "PAID",
		Closed:    "CLOSED",
		ReEntered: "RE_ENTERED",
	}
)

func (o OrderStatus) String() string {
	switch o {
	case Draft:
		return "DRAFT"
	case Entered:
		return "ENTERED"
	case Paid:
		return "PAID"
	case Closed:
		return "CLOSED"
	case ReEntered:
		return "RE_ENTERED"
	default:
		return "invalid"
	}
}

type (

	// BuildModel is an exported
	BuildModel interface {
		Build(args ...interface{}) error
	}

	// Order represents the structure send lio
	Order struct {
		BuildModel   interface{}   `json:"_,omitempty"`
		UUID         string        `json:"id"`
		Number       string        `json:"number,omitempty"`
		Reference    string        `json:"reference,omitempty"`
		Notes        string        `json:"notes,omitempty"`
		CreatedAt    time.Time     `json:"created_at,omitempty"`
		UpdatedAt    time.Time     `json:"updated_at,omitempty"`
		MerchantID   string        `json:"merchant_id" valid:"required"`
		LogicNumber  string        `json:"logic_number" valid:"required"`
		Status       OrderStatus   `json:"status" valid:"required"`
		Ref          string        `json:"ref"`
		SyncCode     int           `json:"sync_code"`
		Items        []Item        `json:"items,omitempty" `
		Transactions []Transaction `json:"transactions,omitempty"`
	}

	// OrderPg exported
	OrderPg struct {
		UUID    string `gorm:"primary_key"`
		Payload JSONB  `gorm:"not null, column:payload"`
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
	o.SyncCode = 200
	if len(strings.TrimSpace(o.UUID)) == 0 {
		o.UUID = uuid.NewV4().String()
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

// Value exported
func (j JSONB) Value() (driver.Value, error) {
	if j.IsNull() {
		//      log.Trace("returning null")
		return nil, nil
	}
	return string(j), nil
}

// Scan exported
func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source was not string")
	}
	// I think I need to make a copy of the bytes.
	// It seems the byte slice passed in is re-used
	*j = append((*j)[0:0], s...)

	return nil
}

// MarshalJSON returns *m as the JSON encoding of m.
func (j JSONB) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (j *JSONB) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

// IsNull exported
func (j JSONB) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}

// Equals exported
func (j JSONB) Equals(j1 JSONB) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}

// UnmarshalJSON sets *m to a copy of data.
func (o *OrderStatus) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("OrderStatus should be a string, got %s", data)
	}

	v, ok := OrderStatusToValue[strings.ToUpper(s)]
	if !ok {
		return fmt.Errorf("invalid OrderStatus %q", s)
	}
	*o = v
	return nil
}

// MarshalJSON returns *o as the JSON encoding of m.
func (o OrderStatus) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(o).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := OrderStatusToName[o]
	if !ok {
		return nil, fmt.Errorf("invalid OrderStatus: %d", o)
	}
	return json.Marshal(s)
}
