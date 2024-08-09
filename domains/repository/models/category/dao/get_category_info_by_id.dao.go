package dao

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetCategoryInfoByIDDAORes struct {
	ID                primitive.ObjectID                 `bson:"_id"`
	Name              string                             `bson:"name"`
	Genre             models.GENRE                       `bson:"genre"`
	TotalParticipants int                                `bson:"total_participants"`
	RangeMovement     models.RANGE_MOVEMENT              `bson:"range_movement"`
	Sport             models.SPORT                       `bson:"sport"`
	Organizer         GetCategoryInfoOrganizerByIDDAORes `bson:"organizer"`
}

type GetCategoryInfoOrganizerByIDDAORes struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
}
