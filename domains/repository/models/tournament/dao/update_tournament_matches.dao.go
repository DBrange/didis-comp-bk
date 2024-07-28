package dao

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateTournamentMatchesDAOReq struct {
	Matches              *[]primitive.ObjectID `bson:"matches,omitempty"`
	common.UpdateBaseDAO `bson:",inline"`
}
