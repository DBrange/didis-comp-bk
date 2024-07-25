package dao

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateGuestCompetitorDAOReq struct {
	GuestPlayerID primitive.ObjectID `bson:"guest_competitor_id"`
	CompetitorID  primitive.ObjectID `bson:"competitor_id"`
common.CreateBaseDAO `bson:",inline"`
}
