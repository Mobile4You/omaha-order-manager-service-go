package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Terminal is an exported
	Terminal struct {
		UUID      bson.ObjectId `json:"id" bson:"_id"`
		Number    string        `json:"number" validate:"required"`
		CreatedAt time.Time     `json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt time.Time     `json:"updated_at,omitempty" bson:"updated_at"`
	}
)
