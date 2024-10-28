package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetUserTournamentsDAORes struct {
	Tournaments []*GetUserTournamentDAORes `bson:"tournaments"`
	Total       int                        `bson:"total"`
}
type GetUserTournamentDAORes struct {
	ID           *primitive.ObjectID             `bson:"_id"`
	Name         string                          `bson:"name"`
	StartDate    *time.Time                      `bson:"start_date"`
	FinishDate   *time.Time                      `bson:"finish_date"`
	Points       int                             `bson:"points"`
	Image        *string                         `bson:"image"`
	AverageScore float32                         `bson:"average_score"`
	TotalPrize   float64                         `bson:"total_prize"`
	Location     *dao.GetLocationByIDDAORes      `bson:"location"`
	Organizer    *GetUserTournamentsOrganizerDAO `bson:"organizer"`
}
