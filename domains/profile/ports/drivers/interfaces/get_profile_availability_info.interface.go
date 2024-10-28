package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type GetProfileDailyAvailabilityByID interface {
	GetProfileDailyAvailabilityByID(ctx context.Context, userID string, day string) (*models.GetDailyAvailabilityByIDDTORes,string, error)
}
