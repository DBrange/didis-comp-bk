package dao

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateRoundDAOReq struct {
	TournamentID         primitive.ObjectID `bson:"tournament_id"`
	Name                 models.ROUND       `bson:"round"`
	TotalPrize           float64            `bson:"total_prize"`
	Points               int                `bson:"points"`
	TotalClassified      int                `bson:"total_classified"`
	BestThird            int                `bson:"best_third"`
	common.CreateBaseDAO `bson:",inline"`
}
