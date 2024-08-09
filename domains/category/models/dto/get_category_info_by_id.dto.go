package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type GetCategoryInfoByIDDTORes struct {
	ID                string                             `bson:"_id"`
	Name              string                             `bson:"name"`
	Genre             models.GENRE                       `bson:"genre"`
	TotalParticipants int                                `bson:"total_participants"`
	RangeMovement     models.RANGE_MOVEMENT              `bson:"range_movement"`
	Sport             models.SPORT                       `bson:"sport"`
	Organizer         GetCategoryInfoOrganizerByIDDTORes `bson:"organizer"`
}

type GetCategoryInfoOrganizerByIDDTORes struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
}
