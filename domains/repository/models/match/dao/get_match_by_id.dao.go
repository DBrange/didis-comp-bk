package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/round/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetMatchDAORes struct {
	ID             *primitive.ObjectID                        `bson:"_id"`
	Date           *time.Time                                 `bson:"date"`
	Result         string                                     `bson:"result"`
	Winner         *primitive.ObjectID                        `bson:"winner"`
	Position       int                                        `bson:"position"`
	PositionWinner *int                                       `bson:"position_winner"`
	Sport          models.SPORT                               `bson:"sport"`
	Competitors    []*dao.GetRoundWithMatchesCompetitorDAORes `bson:"competitors"`
	Tournament     *GetMatchTorurnamentDAORes                 `bson:"tournament"`
	Round          *GetMatchRoundDAORes                       `bson:"round"`
}

type GetMatchTorurnamentDAORes struct {
	ID   *primitive.ObjectID `bson:"_id"`
	Name string              `bson:"name"`
}
type GetMatchRoundDAORes struct {
	ID    *primitive.ObjectID `bson:"_id"`
	Round models.ROUND        `bson:"round"`
}

// type GetMatchByIDDAORes struct {
// 	ID           *primitive.ObjectID `bson:"_id"`
// 	Sport        *models.SPORT       `bson:"sport"`
// 	RoundID      *primitive.ObjectID `bson:"round_id"`
// 	Result       *primitive.ObjectID `bson:"result"`
// 	Winner       *primitive.ObjectID `bson:"winner"`
// 	TournamentID *primitive.ObjectID `bson:"tournament_id"`
// 	Position     int                 `bson:"position"`
// 	Date         *time.Time          `bson:"date"`
// 	// Votes        map[string]string  `bson:"votes"`
// 	CreatedAt time.Time  `bson:"created_at"`
// 	UpdatedAt time.Time  `bson:"updated_at"`
// 	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
// }
