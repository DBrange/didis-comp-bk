package dao

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type GetTournamentFiltersDAORes struct {
	Surface          models.TENNIS_SURFACE  `bson:"surface"`
	Sport            models.SPORT           `bson:"sport"`
	CompetitorType   models.COMPETITOR_TYPE `bson:"competitor_type"`
	MaxCapacity      int                    `bson:"max_capacity"`
	TotalCompetitors int                    `bson:"total_competitors"`
}
