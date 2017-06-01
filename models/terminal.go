package models

import (
	"time"
)

type (
	// Terminal is an exported
	Terminal struct {
		UUID      string    `json:"id"`
		Number    string    `json:"number" validate:"required"`
		CreatedAt time.Time `json:"created_at,omitempty"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
	}
)
