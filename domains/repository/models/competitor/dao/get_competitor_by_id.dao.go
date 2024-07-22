package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetCompetitorByIDDAORes struct {
	ID        *primitive.ObjectID `bson:"_id"`
	SingleID  *primitive.ObjectID `bson:"single_id"`
	DoubleID  *primitive.ObjectID `bson:"double_id"`
	TeamID    *primitive.ObjectID `bson:"team_id"`
	Sport     models.SPORT        `bson:"sport"`
	CreatedAt time.Time           `bson:"created_at"`
	UpdatedAt time.Time           `bson:"updated_at"`
	DeletedAt *time.Time          `bson:"deleted_at,omitempty"`
}
