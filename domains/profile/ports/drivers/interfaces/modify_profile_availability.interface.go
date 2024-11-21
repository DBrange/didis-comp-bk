package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type ModifyProfileAvailability interface {
	ModifyProfileAvailability(ctx context.Context,userID, availabilityID string, availabilityInfoDTO *models.UpdateDailyAvailabilityDTOReq) error
}
