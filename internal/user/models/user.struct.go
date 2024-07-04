package models

import "time"

type User struct {
	ID        string    `json:"id" bson:"_id"`
	FirstName string    `json:"first_name" bson:"first_name" validate:"required"`
	LastName  string    `json:"last_name" bson:"last_name"`
	Role      []ROLE    `json:"role" bson:"role"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
}