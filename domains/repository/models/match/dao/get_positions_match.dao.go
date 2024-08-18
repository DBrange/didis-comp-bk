package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetPositionsBracketMatchDAORes struct {
	ID            *primitive.ObjectID                         `bson:"_id"`
	PositionMatch int                                         `bson:"position_match"`
	Competitors   []*GetPositionsBracketMatchCompetitorDAORes `bson:"competitors"`
}

type GetPositionsBracketMatchCompetitorDAORes struct {
	ID              *primitive.ObjectID `bson:"_id"`
	Position        int                 `bson:"position"`
	CurrentPosition *int                `bson:"current_position"`
}
