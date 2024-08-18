package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetTournamentInfoToFinaliseItDAORes struct {
	CategoryID *primitive.ObjectID `bson:"category_id"`
	TotalPrize float64             `bson:"total_prize"`
	Points     int                 `bson:"points"`
}
