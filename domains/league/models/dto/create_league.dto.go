package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type CreateLeagueDTOReq struct {
	Name              string                `bson:"name"`
	Genre             models.GENRE          `bson:"genre"`
	TotalParticipants int                   `bson:"total_participants"`
	RangeMovement     models.RANGE_MOVEMENT `bson:"range_movement"`
	AverageScore      float32               `bson:"average_score"`
	Sport             models.SPORT          `bson:"sport"`
}
