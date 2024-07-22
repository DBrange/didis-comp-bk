package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type RoleDAOReq struct {
	Name      models.ROLE        `bson:"name"`
	RoleType  models.ROLE_TYPE   `bson:"role_type"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
}
