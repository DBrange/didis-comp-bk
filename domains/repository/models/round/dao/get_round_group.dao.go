package dao

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetRoundGroupsDAORes struct {
	ID              *primitive.ObjectID    `bson:"_id"`
	Round           models.ROUND           `bson:"round"`
	TotalPrize      float64                `bson:"total_prize"`
	Points          int                    `bson:"points"`
	TotalClassified int                    `bson:"total_classified"`
	BestThird       int                    `bson:"best_third"`
	Groups          []*GetRoundGroupDAORes `bson:"groups"`
}

type GetRoundGroupDAORes struct {
	ID          *primitive.ObjectID                       `bson:"_id"`
	Position    int                                       `bson:"position"`
	Matches     []*GetRoundWithMatchesMatchDAORes         `bson:"matches"`
	Competitors []*GetRoundGroupCompetitorWithStatsDAORes `bson:"competitors"`
}

type GetRoundGroupCompetitorWithStatsDAORes struct {
	ID              *primitive.ObjectID                  `bson:"_id"`
	CurrentPosition *int                                 `bson:"current_position"`
	Position        int                                  `bson:"position"`
	Stats           TournamentGroupCompetitorStatsDAOReq `bson:"stats"`
	Users           []*GetRoundWithMatchesUserDAORes     `bson:"users"`
	GuestUsers      []*GetRoundWithMatchesUserDAORes     `bson:"guest_users"`
}

type TournamentGroupCompetitorStatsDAOReq struct {
	MatchesPlayed   int   `bson:"matches_played"`
	MatchesLost     int   `bson:"matches_lost"`
	MatchesWon      int   `bson:"matches_won"`
	SetsWon         int   `bson:"sets_won"`
	SetsLost        int   `bson:"sets_lost"`
	GamesWon        int   `bson:"games_won"`
	GamesLost       int   `bson:"games_lost"`
	LastFiveMatches []int `bson:"last_five_matches"`
}
