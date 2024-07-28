package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserDAOReq struct {
	FirstName  string               `bson:"first_name"`
	LastName   string               `bson:"last_name"`
	Username   *string              `bson:"username"`
	Birthdate  *time.Time           `bson:"birthdate"`
	Password   *string              `bson:"password"`
	Email      string               `bson:"email"`
	Phone      *string              `bson:"phone"`
	Image      *string              `bson:"image"`
	Active     bool                 `bson:"active"`
	Genre      models.GENRE         `bson:"genre"`
	Roles      []primitive.ObjectID `bson:"roles"`
	LocationID *primitive.ObjectID  `bson:"location_id"`
	PaymentID  *primitive.ObjectID  `bson:"payment_id"`
	CreatedAt  time.Time            `bson:"created_at"`
	UpdatedAt  time.Time            `bson:"updated_at"`
	DeletedAt  *time.Time           `bson:"deleted_at,omitempty"`
}

func (u *CreateUserDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}
