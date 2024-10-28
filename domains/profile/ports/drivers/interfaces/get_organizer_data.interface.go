package interfaces

import (
	"context"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type GetOrganizerData interface {
	GetOrganizerData(ctx context.Context, userID string) (*dto.GetOrganizerDataDTORes, error)
}
