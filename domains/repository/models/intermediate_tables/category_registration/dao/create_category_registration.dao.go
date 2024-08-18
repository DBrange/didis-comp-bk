package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateCategoryRegistrationDAOReq struct {
	CompetitorID         primitive.ObjectID       `bson:"competitor_id"`
	CategoryID           primitive.ObjectID       `bson:"category_id"`
	Points               int                      `bson:"points"`
	RegisteredPositions  []RegistedPositionDAORes `bson:"registered_positions"`
	CurrentPosition      *int                     `bson:"current_position"`
	common.CreateBaseDAO `bson:",inline"`
}

type RegistedPositionDAORes struct {
	Date     time.Time `bson:"date"`
	Position int       `bson:"position"`
}
