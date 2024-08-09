package dao

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateCategoryRegistrationDAOReq struct {
	CompetitorID         primitive.ObjectID `bson:"competitor_id"`
	CategoryID           primitive.ObjectID `bson:"category_id"`
	Points               int                `bson:"points"`
	RegisteredPositions  []int              `bson:"registered_positions"`
	CurrentPosition      *int               `bson:"current_position"`
	common.CreateBaseDAO `bson:",inline"`
}
