package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	tournament_registration_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/tournament_registration/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetProfileUserTournamentsDAORes struct {
	Tournaments []*GetProfileUserTournamentDAORes `bson:"tournaments"`
	Total       int                               `bson:"total"`
}
type GetProfileUserTournamentDAORes struct {
	ID        *primitive.ObjectID                    `bson:"_id"`
	Name      string                                 `bson:"name"`
	Location  *dao.GetLocationByIDDAORes             `bson:"location"`
	Organizer *GetUserTournamentsOrganizerDAO        `bson:"organizer"`
	Matches   []*GetProfileUserTournamentMatchDAORes `bson:"matches"`
}

type GetProfileUserTournamentMatchDAORes struct {
	ID          *primitive.ObjectID                                             `bson:"_id"`
	Result      string                                                          `bson:"result"`
	Winner      *primitive.ObjectID                                             `bson:"winner"`
	Date        *time.Time                                                      `bson:"date"`
	Round       *GetProfileUserTournamentRoundDAORes                            `bson:"round"`
	Competitors []*tournament_registration_dao.GetCompetitorsInTournamentDAORes `bson:"competitors"`
}

type GetProfileUserTournamentRoundDAORes struct {
	ID    *primitive.ObjectID `bson:"_id"`
	Round models.ROUND        `bson:"round"`
}
