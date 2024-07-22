package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type ModifyProfileAvailability interface {
	ModifyProfileAvailability(ctx context.Context, availabilityID string, availabilityInfoDTO *dto.ModifyProfileDailyAvailabilityDTOReq) error
}
