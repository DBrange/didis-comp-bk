package dto

import (
	"time"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
)

type GetCategoryRegistrationByIDDTORes struct {
	ID                  string                    `json:"_id"`
	CompetitorID        string                    `json:"competitor_id"`
	CategoryID          string                    `json:"category_id"`
	Points              int                       `json:"points"`
	RegisteredPositions []*RegistedPositionDTORes `json:"registered_positions"`
	CurrentPosition     *int                      `json:"current_position"`
	common.GetBaseDAO   `json:",inline"`
}

type RegistedPositionDTORes struct {
	Date     time.Time `bson:"date"`
	Position int       `bson:"position"`
}
