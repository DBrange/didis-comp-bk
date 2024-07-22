package dao

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateCompetitorUserDAOReq struct {
	UserID       primitive.ObjectID `bson:"user_id"`
	CompetitorID primitive.ObjectID `bson:"competitor_id"`
	common.CreateBaseDAO
}
