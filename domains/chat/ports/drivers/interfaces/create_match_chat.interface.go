package interfaces

import (
	"context"
)

type CreateMatchChat interface {
	CreateMatchChat(ctx context.Context, matchID string, competitorIDs []string, userID string) error
}
