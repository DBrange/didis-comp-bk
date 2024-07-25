package dao

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetLeagueRegistrationByIDDAORes struct {
	ID                  primitive.ObjectID `bson:"_id"`
	CompetitorID        primitive.ObjectID `bson:"competitor_id"`
	LeagueID            primitive.ObjectID `bson:"league_id"`
	Points              int                `bson:"points"`
	RegisteredPositions []int              `bson:"registered_positions"`
	CurrentPosition     *int               `bson:"registered_positions"`
	common.GetBaseDAO `bson:",inline"`
}
