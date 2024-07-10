package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	user_coll     *mongo.Collection
	location_coll *mongo.Collection
}

func NewRepository(user_coll *mongo.Collection, location_coll *mongo.Collection) *Repository {
	return &Repository{
		user_coll:     user_coll,
		location_coll: location_coll,
	}
}
