package dao

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateCompetitorMatchDAOReq struct {
	MatchID              *primitive.ObjectID `bson:"match_id,omitempty"`
	CompetitorID         *primitive.ObjectID `bson:"competitor_id,omitempty"`
	Position             *int                `bson:"position,omitempty"`
	common.UpdateBaseDAO `bson:",inline"`
}
