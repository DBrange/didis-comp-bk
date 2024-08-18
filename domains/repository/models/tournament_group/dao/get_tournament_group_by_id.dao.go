package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetTournamentGroupDAORes struct {
	ID           *primitive.ObjectID                `bson:"_id"`
	TournamentID primitive.ObjectID                 `bson:"tournament_id"`
	Competitors  []*TournamentGroupCompetitorDAOReq `bson:"competitors"`
	Matches      []*primitive.ObjectID              `bson:"matches"`
	Position     int                                `bson:"position"`
}
