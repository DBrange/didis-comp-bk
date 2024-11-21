package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type GetCategoryInfoByIDDTORes struct {
	ID                string                             `json:"id"`
	Name              string                             `json:"name"`
	Genre             models.GENRE                       `json:"genre"`
	TotalParticipants int                                `json:"total_participants"`
	RangeMovement     models.RANGE_MOVEMENT              `json:"range_movement"`
	Sport             models.SPORT                       `json:"sport"`
	CompetitorType    models.COMPETITOR_TYPE             `json:"competitor_type"`
	Organizer         GetCategoryInfoOrganizerByIDDTORes `json:"organizer"`
}

type GetCategoryInfoOrganizerByIDDTORes struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
