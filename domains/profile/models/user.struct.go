package models

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	ID         string               `json:"id" bson:"_id"`
	FirstName  string               `json:"first_name" bson:"first_name"`
	LastName   string               `json:"last_name" bson:"last_name"`
	Username   *string              `json:"username" bson:"username"`
	Birthdate  *time.Time           `json:"birthdate" bson:"birthdate"`
	Password   *string              `json:"password" bson:"password"`
	Email      string               `json:"email" bson:"email"`
	Phone      *string              `json:"phone" bson:"phone"`
	Image      *string              `json:"image" bson:"image"`
	Active     bool                 `json:"active" bson:"active"`
	Genre      models.GENRE         `json:"genre" bson:"genre"`
	Roles      []primitive.ObjectID `json:"role" bson:"role"`
	LocationID *string              `json:"location_id" bson:"location_id"`
	PaymentID  *string              `json:"payment_id" bson:"payment_id"`
	CreatedAt  time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at" bson:"updated_at"`
	DeletedAt  *time.Time           `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}
