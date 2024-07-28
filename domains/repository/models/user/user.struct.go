package models

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         string               `bson:"_id"`
	FirstName  string               `bson:"first_name"`
	LastName   string               `bson:"last_name"`
	Username   string               `bson:"username"`
	Birthdate  int8                 `bson:"birthdate"`
	Password   string               `bson:"password"`
	Email      string               `bson:"email"`
	Phone      string               `bson:"phone"`
	Image      string               `bson:"image"`
	Active     bool                 `bson:"active"`
	Genre      models.GENRE         `bson:"genre"`
	Roles      []primitive.ObjectID `bson:"roles"`
	LocationID string               `bson:"location_id"`
	PaymentID  string               `bson:"payment_id"`
	CreatedAt  time.Time            `bson:"created_at"`
	UpdatedAt  time.Time            `bson:"updated_at"`
	DeletedAt  *time.Time           `bson:"deleted_at,omitempty"`
}
