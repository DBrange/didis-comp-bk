package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetTournamentsFromCategoryDAORes struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Points      *int               `bson:"points"`
	StartDate   *string            `bson:"start_date"`
	FinishtDate *string            `bson:"finish_date"`
}
