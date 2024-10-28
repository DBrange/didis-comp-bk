package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetUserCategoriesDAO struct {
	Categories []*GetUserCategoriesCategoryDAO `bson:"categories"`
}

type GetUserCategoriesCategoryDAO struct {
	ID                 *primitive.ObjectID                 `bson:"_id"`
	Name               string                              `bson:"name"`
	CompetitorData     *GetUserCategoriesCompetitorDataDAO `bson:"competitor_data"`
	Organizer          *GetUserCategoriesOrganizerDAO      `bson:"organizer"`
}

type GetUserCategoriesOrganizerDAO struct {
	ID        *primitive.ObjectID `bson:"_id"`
	FirstName string              `bson:"first_name"`
	LastName  string              `bson:"last_name"`
}

type GetUserCategoriesCompetitorDataDAO struct {
	Points          int `bson:"points"`
	CurrentPosition int `bson:"current_position"`
	Users []*GetUserCategoriesUserDAO `bson:"users"`
}

type GetUserCategoriesUserDAO struct{
	ID        *primitive.ObjectID `bson:"_id"`
	FirstName string              `bson:"first_name"`
	LastName  string              `bson:"last_name"`
}