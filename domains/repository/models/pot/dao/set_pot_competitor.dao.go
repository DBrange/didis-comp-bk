package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type SetPotCompetitorDAOReq struct {
	PotID       *primitive.ObjectID   `bson:"pot_id"`
	Competitors []*primitive.ObjectID `bson:"competitors"`
}
