package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetUserPrimaryDataDAORes struct {
	ID        *primitive.ObjectID `bson:"_id"`
	FirstName string              `bson:"first_name"`
	LastName  string              `bson:"last_name"`
	Username  string              `bson:"username"`
	Image     string              `bson:"image"`
}
