package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetGuestPlayerByIDDAORes struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Email     string             `bson:"email"`
	Image     *string            `bson:"image"`
	Genre     models.GENRE       `bson:"genre"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
}
