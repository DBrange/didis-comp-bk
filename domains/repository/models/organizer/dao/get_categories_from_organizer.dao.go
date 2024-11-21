package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetCategoriesFromOrganizerDAORes struct {
	CategoryID        primitive.ObjectID                           `bson:"_id"`
	Name              string                                       `bson:"name"`
	Competitors       []GetCategoriesFromOrganizerCompetitorDAORes `bson:"competitors"`
	TotalParticipants int32                                        `bson:"total"`
}

type GetCategoriesFromOrganizerCompetitorDAORes struct {
	CompetitorID        primitive.ObjectID                     `bson:"_id"`
	CurrentPosition     *int                                   `bson:"current_position"`
	RegisteredPositions []RegisteredPositionsDAORes            `bson:"registered_positions"`
	Points              int                                    `bson:"points"`
	Users               []GetCategoriesFromOrganizerUserDAORes `bson:"users"`
}

type GetCategoriesFromOrganizerUserDAORes struct {
	UserID    primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Image     *string            `bson:"image"`
}

type RegisteredPositionsDAORes struct {
	Date     time.Time `bson:"date"`
	Position int       `bson:"position"`
}
