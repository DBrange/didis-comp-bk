package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type CreateCategoryDTOReq struct {
	Name              string                 `json:"name"`
	Genre             models.GENRE           `json:"genre"`
	TotalParticipants int                    `json:"total_participants"`
	RangeMovement     models.RANGE_MOVEMENT  `json:"range_movement"`
	AverageScore      float32                `json:"average_score"`
	Sport             models.SPORT           `json:"sport"`
	CompetitorType    models.COMPETITOR_TYPE `json:"competitor_type"`
}
