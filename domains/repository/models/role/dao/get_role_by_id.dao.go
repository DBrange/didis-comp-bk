package dao

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetRoleDAOByID struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     models.ROLE        `bson:"name"`
	RoleType models.ROLE_TYPE   `bson:"role_type"`
}
