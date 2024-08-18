package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type AddCompetitorsToTournamentGroupsDAOReq struct {
	GroupID     *primitive.ObjectID `bson:"group_id"`
	Competitors []*primitive.ObjectID `bson:"competitors"`
}
