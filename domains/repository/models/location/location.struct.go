package models

import "time"

type Location struct {
	State     *string    `bson:"state"`
	Country   *string    `bson:"country"`
	City      *string    `bson:"city"`
	Lat       *string    `bson:"lat"`
	Long      *string    `bson:"long"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
