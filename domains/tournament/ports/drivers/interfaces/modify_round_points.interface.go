package interfaces

import "context"

type ModifyRoundPoints interface {
	ModifyRoundPoints(ctx context.Context, roundID string, points int) error
}
