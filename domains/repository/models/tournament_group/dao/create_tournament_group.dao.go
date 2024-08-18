package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTournamentGroupDAOReq struct {
	TournamentID primitive.ObjectID                      `bson:"tournament_id"`
	Competitors  []*TournamentGroupCompetitorDAOReq `bson:"competitors"`
	Matches      []*primitive.ObjectID                   `bson:"matches"`
	Position     int                                     `bson:"position"`
	CreatedAt    time.Time                               `bson:"created_at"`
	UpdatedAt    time.Time                               `bson:"updated_at"`
	DeletedAt    *time.Time                              `bson:"deleted_at,omitempty"`
}

type TournamentGroupCompetitorDAOReq struct {
	CompetitorID *primitive.ObjectID                  `bson:"competitor_id"`
	Stats        TournamentGroupCompetitorStatsDAOReq `bson:"stats"`
}
type TournamentGroupCompetitorStatsDAOReq struct {
	MatchesPlayed int `bson:"matches_played"`
	MatchesLost   int `bson:"matches_lost"`
	MatchesWon    int `bson:"matches_won"`
	SetsWon       int `bson:"sets_won"`
	SetsLost      int `bson:"sets_lost"`
	GamesWon      int `bson:"games_won"`
	GamesLost     int `bson:"games_lost"`
}

func (u *CreateTournamentGroupDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}
