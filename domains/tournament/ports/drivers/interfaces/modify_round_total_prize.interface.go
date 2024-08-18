package interfaces

import "context"

type ModifyRoundTotalPrize interface {
	ModifyRoundTotalPrize(ctx context.Context, roundID string, totalPrize float64) error
}
