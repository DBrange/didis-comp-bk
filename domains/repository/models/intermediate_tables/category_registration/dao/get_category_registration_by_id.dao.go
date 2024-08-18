package dao

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetCategoryRegistrationByIDDAORes struct {
	ID                  *primitive.ObjectID       `bson:"_id"`
	CompetitorID        *primitive.ObjectID       `bson:"competitor_id"`
	CategoryID          *primitive.ObjectID       `bson:"category_id"`
	Points              int                      `bson:"points"`
	RegisteredPositions []*RegistedPositionDAORes `bson:"registered_positions"`
	CurrentPosition     *int                     `bson:"current_position"`
	common.GetBaseDAO   `bson:",inline"`
}

