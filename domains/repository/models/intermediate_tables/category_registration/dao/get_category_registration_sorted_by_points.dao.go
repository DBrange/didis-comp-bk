package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetCategoryRegistrationSortedByPointsDAORes struct {
	CompetitorID    *primitive.ObjectID `bson:"competitor_id"`
	CurrentPosition *int                `bson:"current_position"`
}
