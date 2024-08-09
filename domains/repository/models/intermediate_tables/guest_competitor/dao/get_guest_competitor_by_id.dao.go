package dao

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetGuestCompetitorByIDDAORes struct {
	GuestUserID       primitive.ObjectID `bson:"guest_user_id"`
	CompetitorID      primitive.ObjectID `bson:"competitor_id"`
	common.GetBaseDAO `bson:",inline"`
}
