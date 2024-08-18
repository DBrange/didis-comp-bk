package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetRoundWithCompetitorsDAORes struct {
	ID            *primitive.ObjectID   `bson:"_id"`
	TotalPrize    float64               `bson:"total_prize"`
	Points        int                   `bson:"points"`
	CompetitorIDs []*primitive.ObjectID `bson:"competitor_ids"`
}
