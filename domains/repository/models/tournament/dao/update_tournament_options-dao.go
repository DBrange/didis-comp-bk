package dao

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateTournamentOptionsDAOReq struct {
	DoubleEliminationID  *primitive.ObjectID   `bson:"double_elimination_id,omitempty"`
	Rounds               *[]primitive.ObjectID `bson:"rounds,omitempty"`
	Pots                 *[]primitive.ObjectID `bson:"pots,omitempty"`
	Groups               *[]primitive.ObjectID `bson:"groups,omitempty"`
	common.UpdateBaseDAO `bson:",inline"`
}
