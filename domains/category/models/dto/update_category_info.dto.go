package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type UpdateCategoryDTOReq struct {
	Name *string `json:"name,omitempty"`
	// Genre             *models.GENRE          `json:"genre,omitempty"`
	TotalParticipants *int                   `json:"total_participants,omitempty"`
	RangeMovement *models.RANGE_MOVEMENT `json:"range_movement,omitempty" validate:"rangeMovement,omitempty"`
	// AverageScore      *float32               `json:"average_score,omitempty"`
	// Tournaments       *[]primitive.ObjectID  `json:"tournaments,omitempty"`
}
