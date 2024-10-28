package interfaces

import (
	"context"
	"time"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type GetUserFollowers interface {
	GetUserFollowers(ctx context.Context, userID string,name string, limit int, lastCreatedAt *time.Time) (*dto.GetUserFollowersDTORes, error)
}
