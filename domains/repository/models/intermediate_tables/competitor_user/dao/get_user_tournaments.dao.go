package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetUserTournamentsDAO struct {
	Tournaments []*GetUserTournamentsTournamentDAO `bson:"tournaments"`
}

type GetUserTournamentsTournamentDAO struct {
	ID           *primitive.ObjectID             `bson:"_id"`
	Name         string                          `bson:"name"`
	StartDate    *time.Time                      `bson:"start_date"`
	FinishDate   *time.Time                      `bson:"finish_date"`
	Points       int                             `bson:"points"`
	AverageScore float32                         `bson:"average_score"`
	Location     *dao.GetLocationByIDDAORes      `bson:"location"`
	Matches      []*primitive.ObjectID           `bson:"matches"`
	Organizer    *GetUserTournamentsOrganizerDAO `bson:"organizer"`
}

type GetUserTournamentsOrganizerDAO struct {
	ID        *primitive.ObjectID `bson:"_id"`
	FirstName string              `bson:"first_name"`
	LastName  string              `bson:"last_name"`
}

// type GetUserTournamentsDAO struct {
// 	Tournaments []*GetUserTournamentsTournamentDAO `bson:"tournaments"`
// }

// type GetUserTournamentsTournamentDAO struct {
// 	ID           *primitive.ObjectID             `bson:"_id"`
// 	Name         string                          `bson:"name"`
// 	StartDate    *time.Time                      `bson:"start_date"`
// 	FinishDate   *time.Time                      `bson:"finish_date"`
// 	Points       int                             `bson:"points"`
// 	AverageScore float32                         `bson:"average_score"`
// 	Location     *dao.GetLocationByIDDAORes      `bson:"location"`
// 	Matches      []*primitive.ObjectID   `bson:"matches"`
// 	Organizer    *GetUserTournamentsOrganizerDAO `bson:"organizer"`
// }

// type GetUserTournamentsOrganizerDAO struct {
// 	ID        *primitive.ObjectID `bson:"_id"`
// 	FirstName string              `bson:"first_name"`
// 	LastName  string              `bson:"last_name"`
// }

// type GetUserTournamentsMatchDAO struct {
// 	ID              *primitive.ObjectID                     `bson:"_id"`
// 	Result          string                                  `bson:"result"`
// 	Winner          *primitive.ObjectID                     `bson:"winner"`
// 	Round           *GetUserTournamentsRoundDAO             `bson:"round"`
// 	Competitors []*GetUserTournamentsRivalCompetitorDAO `bson:"competitors"`
// 	RivalCompetitors []*GetUserTournamentsRivalCompetitorDAO `bson:"rival_competitors"`
// }

// type GetUserTournamentsRoundDAO struct {
// 	ID    *primitive.ObjectID          `bson:"_id"`
// 	Round []*GetUserTournamentsUserDAO `bson:"round"`
// }

// type GetUserTournamentsRivalCompetitorDAO struct {
// 	ID    *primitive.ObjectID          `bson:"_id"`
// 	Users []*GetUserTournamentsUserDAO `bson:"users"`
// }

// type GetUserTournamentsUserDAO struct {
// 	ID        *primitive.ObjectID `bson:"_id"`
// 	FirstName string              `bson:"first_name"`
// 	LastName  string              `bson:"last_name"`
// }
