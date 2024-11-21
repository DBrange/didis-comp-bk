package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetCompetitorsOfCategoryDAORes struct {
	Total       int64                                       `bson:"total"`
	Competitors []*GetCompetitorsOfCategoryCompetitorDAORes `bson:"competitors"`
}
type GetCompetitorsOfCategoryCompetitorDAORes struct {
	ID                  primitive.ObjectID                    `bson:"_id"`
	CurrentPosition     *int                                  `bson:"current_position"`
	RegisteredPositions []RegistedPositionDAORes              `bson:"registered_positions"`
	Points              int                                   `bson:"points"`
	Users               []*GetCompetitorsOfCategoryUserDAORes `bson:"users"`
	GuestUsers          []*GetCompetitorsOfCategoryUserDAORes `bson:"guest_users"`
}

type GetCompetitorsOfCategoryUserDAORes struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Image     string             `bson:"image"`
	Username     string             `bson:"username"`
}
