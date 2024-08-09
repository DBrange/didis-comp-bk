package dao

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateGuestCompetitorDAOReq struct {
	GuestUserID          primitive.ObjectID `bson:"guest_user_id"`
	CompetitorID         primitive.ObjectID `bson:"competitor_id"`
	common.CreateBaseDAO `bson:",inline"`
}
