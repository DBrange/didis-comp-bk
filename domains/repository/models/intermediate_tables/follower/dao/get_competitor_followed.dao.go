package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetCompetitorFollowedDAORes struct {
	ID    *primitive.ObjectID                `bson:"_id"`
	Users []*GetUserCompetitorFollowedDAORes `bson:"users"`
}

type GetUserCompetitorFollowedDAORes struct {
	ID        *primitive.ObjectID `bson:"_id"`
	FirstName string              `bson:"first_name"`
	LastName  string              `bson:"last_name"`
	Image     string              `bson:"image"`
}