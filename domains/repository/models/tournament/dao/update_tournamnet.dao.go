package dao

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
)

type UpdateTournamentInfoDAOReq struct {
	Name                 *string                     `bson:"name,omitempty"`
	Points               *int                        `bson:"points,omitempty"`
	TotalPrize           *float64                    `bson:"total_prize,omitempty"`
	TotalCompetitors     *models.TOURNAMENT_CAPACITY `bson:"total_competitors,omitempty"`
	AverageScore         *float32                    `bson:"average_score,omitempty"`
	Surface              *models.TENNIS_SURFACE      `bson:"surface,omitempty"`
	common.UpdateBaseDAO `bson:",inline"`
}
