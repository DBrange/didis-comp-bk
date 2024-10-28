package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetLocationByIDDAORes struct {
	ID      *primitive.ObjectID `bson:"_id"`
	State   *string      `bson:"state"`
	Country *string      `bson:"country"`
	City    *string      `bson:"city"`
	Lat     *string      `bson:"lat"`
	Long    *string      `bson:"long"`
}
