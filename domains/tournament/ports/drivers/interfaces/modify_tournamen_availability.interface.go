package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type ModifyTournamentAvailability interface {
ModifyTournamentAvailability(ctx context.Context, availabilityID string, availabilityInfoDTO *models.UpdateDailyAvailabilityDTOReq) error 
}
