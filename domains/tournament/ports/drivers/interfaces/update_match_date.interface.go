package interfaces

import (
	"context"
	"time"
)

type UpdateMatchDate interface {
	UpdateMatchDate(ctx context.Context, matchID string, date *time.Time) error
}
