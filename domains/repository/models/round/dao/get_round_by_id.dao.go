package dao

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetRoundByIDDAORes struct {
	ID           primitive.ObjectID `bson:"_id"`
	TournamentID primitive.ObjectID `bson:"tournament_id"`
	Name         models.ROUND       `bson:"round"`
	TotalPrize   float64            `bson:"total_prize"`
}