package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetCategoriesFromOrganizerDAORes struct {
	CategoryID        primitive.ObjectID                           `bson:"category_id"`
	Competitors       []GetCategoriesFromOrganizerCompetitorDAORes `bson:"competitors"`
	TotalParticipants int32                                         `bson:"total_participants"`
}

type GetCategoriesFromOrganizerCompetitorDAORes struct {
	CompetitorID        primitive.ObjectID                     `bson:"competitor_id"`
	CurrentPosition     *int                                   `bson:"current_position"`
	RegisteredPositions []int                                  `bson:"registered_positions"`
	Points              int                                    `bson:"points"`
	Users               []GetCategoriesFromOrganizerUserDAORes `bson:"users"`
}

type GetCategoriesFromOrganizerUserDAORes struct {
	UserID    primitive.ObjectID `bson:"user_id"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Image     string             `bson:"image"`
}
