package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type OrganizeCategoryDTOReq struct {
	Name              string                `json:"name"`
	Genre             models.GENRE          `json:"genre"`
	TotalParticipants int                   `json:"total_participants"`
	RangeMovement     models.RANGE_MOVEMENT `json:"range_movement"`
	Sport             models.SPORT          `json:"sport"`
}
