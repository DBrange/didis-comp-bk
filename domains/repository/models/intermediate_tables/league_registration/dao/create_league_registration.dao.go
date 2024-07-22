package dao

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateLeagueRegistrationDAOReq struct {
	CompetitorID        primitive.ObjectID `bson:"competitor_id"`
	LeagueID            primitive.ObjectID `bson:"league_id"`
	Points              int                `bson:"points"`
	RegisteredPositions []int              `bson:"registered_positions"`
	CurrentPosition     *int               `bson:"registered_positions"`
	common.CreateBaseDAO
}
