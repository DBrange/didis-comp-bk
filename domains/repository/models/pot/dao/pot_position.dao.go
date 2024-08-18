package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type PotOrGroupPositionDAORes struct {
	ID       *primitive.ObjectID `bson:"_id"`
	Position int                 `bson:"position"`
}
