package dao

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateTournamentDAOReq struct {
	Name                *string                `bson:"name,omitempty"`
	Points              *int                   `bson:"points,omitempty"`
	TotalPrize          *float64               `bson:"total_prize,omitempty"`
	TotalCompetitors    *int                   `bson:"total_competitors,omitempty"`
	AverageScore        *float32               `bson:"average_score,omitempty"`
	Surface             *models.TENNIS_SURFACE `bson:"surface,omitempty"`
	LeagueID            *primitive.ObjectID    `bson:"league_id,omitempty"`
	DoubleEliminationID *primitive.ObjectID    `bson:"double_elimination_id,omitempty"`
	Rounds              *[]primitive.ObjectID  `bson:"rounds,omitempty"`
	Matches             *[]primitive.ObjectID  `bson:"matches,omitempty"`
	Pots                *[]primitive.ObjectID  `bson:"pots,omitempty"`
	Groups              *[]primitive.ObjectID  `bson:"groups,omitempty"`
	common.UpdateBaseDAO
}
