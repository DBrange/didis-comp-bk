package models

import "go.mongodb.org/mongo-driver/mongo"

type ValuesForDelete struct {
	MC   *mongo.Collection
	ID   string
	Name string
}
