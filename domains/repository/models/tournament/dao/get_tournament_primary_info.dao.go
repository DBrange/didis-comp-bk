package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_user/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetTournamentPrimaryInfoDAORes struct {
	ID               *primitive.ObjectID                     `bson:"_id"`
	Name             string                                  `bson:"name"`
	FinishDate       *time.Time                              `bson:"finish_date"`
	StartDate        *time.Time                              `bson:"start_date"`
	Points           int                                     `bson:"points"`
	Image            *string                                 `bson:"image"`
	TotalPrize       float64                                 `bson:"total_prize"`
	TotalCompetitors int                                     `bson:"total_competitors"`
	MaxCapacity      models.TOURNAMENT_CAPACITY              `bson:"max_capacity"`
	AverageScore     float32                                 `bson:"average_score"`
	Genre            models.GENRE                            `bson:"genre"`
	Sport            models.SPORT                            `bson:"sport"`
	CompetitorType   models.SPORT                            `bson:"competitor_type"`
	Surface          models.TENNIS_SURFACE                   `bson:"surface"`
	Rounds           []*GetTournamentPrimaryInforRoundDAORes `bson:"rounds"`
	Location         *location_dao.GetLocationByIDDAORes     `bson:"location"`
	Organizer        *dao.GetUserTournamentsOrganizerDAO     `bson:"organizer"`
	Category         *GetTournamentPrimaryInfoCategoryDAORes `bson:"category"`
}

type GetTournamentPrimaryInforRoundDAORes struct {
	ID    *primitive.ObjectID `bson:"_id"`
	Round models.ROUND        `bson:"round"`
}

type GetTournamentPrimaryInfoCategoryDAORes struct {
	ID   *primitive.ObjectID `bson:"_id"`
	Name string              `bson:"name"`
}

// type GetTournamentByIDDAORes struct {
// 	ID               primitive.ObjectID         `bson:"_id"`
// 	Name             string                     `bson:"name"`
// 	FinishDate       *time.Time                 `bson:"finish_date"`
// 	StartDate        *time.Time                 `bson:"start_date"`
// 	Points           *int                       `bson:"points"`
// 	TotalPrize       float64                    `bson:"total_prize"`
// 	TotalCompetitors int                        `bson:"total_competitors"`
// 	MaxCapacity      models.TOURNAMENT_CAPACITY `bson:"max_capacity"`
// 	AverageScore     float32                   `bson:"average_score"`
// 	Availability     TournamentAvailabilityDAO  `bson:"availability"`
// 	Genre               models.GENRE           `bson:"genre"`
// 	Sport               models.SPORT           `bson:"sport"`
// 	Surface             models.TENNIS_SURFACE  `bson:"surface"`
// 	CompetitorType      models.COMPETITOR_TYPE `bson:"competitor_type"`
// 	LocationID          primitive.ObjectID     `bson:"location_id"`
// 	OrganizerID         primitive.ObjectID     `bson:"organizer_id"`
// 	CategoryID          *primitive.ObjectID    `bson:"category_id"`
// 	DoubleEliminationID *primitive.ObjectID    `bson:"double_elimination_id"`
// 	Rounds              []primitive.ObjectID   `bson:"rounds"`
// 	Matches             []primitive.ObjectID   `bson:"matches"`
// 	Pots                []primitive.ObjectID   `bson:"pots"`
// 	Groups              []primitive.ObjectID   `bson:"groups"`
// 	CreatedAt           time.Time              `bson:"created_at"`
// 	UpdatedAt           time.Time              `bson:"updated_at"`
// 	DeletedAt           *time.Time             `bson:"deleted_at,omitempty"`
// }
