package models

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      models.ROLE        `bson:"name"`
	RoleType  models.ROLE_TYPE   `bson:"role_type"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
}
