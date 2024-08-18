package dao

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetRoundWithMatchesDAORes struct {
	ID         *primitive.ObjectID               `bson:"_id"`
	Round      models.ROUND                      `bson:"round"`
	TotalPrize float64                           `bson:"total_prize"`
	Points     int                               `bson:"points"`
	Matches    []*GetRoundWithMatchesMatchDAORes `bson:"matches"`
}

type GetRoundWithMatchesMatchDAORes struct {
	ID             *primitive.ObjectID                    `bson:"_id"`
	Result         string                                 `bson:"result"`
	Winner         *primitive.ObjectID                    `bson:"winner"`
	PositionWinner *int                                   `bson:"position_winner"`
	Competitors    []*GetRoundWithMatchesCompetitorDAORes `bson:"competitors"`
}

type GetRoundWithMatchesCompetitorDAORes struct {
	ID              *primitive.ObjectID              `bson:"_id"`
	CurrentPosition *int                             `bson:"current_position"`
	Position        int                              `bson:"position"`
	Users           []*GetRoundWithMatchesUserDAORes `bson:"users"`
	GuestUsers      []*GetRoundWithMatchesUserDAORes `bson:"guest_users"`
}

type GetRoundWithMatchesUserDAORes struct {
	ID        *primitive.ObjectID `bson:"_id"`
	FirstName string              `bson:"first_name"`
	LastName  string              `bson:"last_name"`
	Image     string              `bson:"image"`
}

// MatchWithUsers representa un match con sus usuarios asociados
// type MatchWithUsers struct {
//     Result     string                          `bson:"result"`
//     Users      []*GetRoundWithMatchesUserDAORes `bson:"users,omitempty"`
//     GuestUsers []*GetRoundWithMatchesUserDAORes `bson:"guest_users,omitempty"`
// }

type RoundWithMatches struct {
	ID         *primitive.ObjectID `bson:"_id"`
	Round      string              `bson:"round"`
	TotalPrize float64             `bson:"total_prize"`
	Matches    []*Match            `bson:"matches"`
}

// Match representa un match individual
type Match struct {
	ID     *primitive.ObjectID `bson:"_id"`
	Result string              `bson:"result"`
}

// CompetitorMatch representa la relación entre un competidor y un match
type CompetitorMatch struct {
	ID           *primitive.ObjectID `bson:"_id"`
	MatchID      *primitive.ObjectID `bson:"match_id"`
	CompetitorID *primitive.ObjectID `bson:"competitor_id"`
	Position     int                 `bson:"position"`
}

// UserWithPosition representa un usuario con su posición en un match específico
type UserWithPosition struct {
	MatchID      *primitive.ObjectID `bson:"match_id"`
	CompetitorID *primitive.ObjectID `bson:"competitor_id"`
	Position     int                 `bson:"position"`
	User         User                `bson:"user"`
}

// User representa la información básica de un usuario
type User struct {
	ID        *primitive.ObjectID `bson:"_id"`
	FirstName string              `bson:"first_name"`
	LastName  string              `bson:"last_name"`
	Image     string              `bson:"image"`
}
