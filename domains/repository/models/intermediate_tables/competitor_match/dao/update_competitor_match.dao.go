package dao

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateCompetitorMatchDAOReq struct {
	MatchID              primitive.ObjectID `bson:"match_id"`
	CompetitorID         primitive.ObjectID `bson:"competitor_id"`
	Position             int                `bson:"position"`
	common.UpdateBaseDAO `bson:",inline"`
}
