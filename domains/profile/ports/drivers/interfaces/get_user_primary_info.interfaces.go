package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type GetUserPrimaryInfo interface {
	GetUserPrimaryInfo(ctx context.Context, fromID, userToID string) (*dto.GetUserPrimatyInfoDTORes, error)
}
