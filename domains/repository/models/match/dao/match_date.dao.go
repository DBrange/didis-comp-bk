package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MatchDateDAOReq struct {
	ID   *primitive.ObjectID `bson:"_id"`
	Date *time.Time          `bson:"date"`
}
