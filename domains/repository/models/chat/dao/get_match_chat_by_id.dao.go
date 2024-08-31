package dao

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetMatchChatByIDDAORes struct {
	ID                 *primitive.ObjectID                  `bson:"_id"`
	AvailabilityStatus models.CHAT_AVAILABILITY_STATUS      `bson:"availability_status"`
	MatchID            *primitive.ObjectID                  `bson:"match_id"`
	Users              []*GetMatchChatByIDUserdDAORes       `bson:"users"`
	Competitors        []*GetMatchChatByIDCompetitordDAORes `bson:"competitors"`
}

type GetMatchChatByIDCompetitordDAORes struct {
	ID                 *primitive.ObjectID             `bson:"_id"`
	AvailabilityStatus models.CHAT_AVAILABILITY_STATUS `bson:"availability_status"`
	Users              []*GetMatchChatByIDUserdDAORes  `bson:"users"`
}

type GetMatchChatByIDUserdDAORes struct {
	ID        *primitive.ObjectID `bson:"_id"`
	FirstName string              `bson:"first_name"`
	LastName  string              `bson:"last_name"`
	Image     string              `bson:"image"`
}
