package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type GetUserByIDDAORes struct {
	ID         string       `bson:"_id"`
	FirstName  string       `bson:"first_name"`
	LastName   string       `bson:"last_name"`
	Username   *string      `bson:"username"`
	Birthdate  *time.Time   `bson:"birthdate"`
	Password   *string      `bson:"password"`
	Email      string       `bson:"email"`
	Phone      *string      `bson:"phone"`
	Image      *string      `bson:"image"`
	Genre      models.GENRE `bson:"genre"`
	LocationID string      `bson:"location_id"`
}
