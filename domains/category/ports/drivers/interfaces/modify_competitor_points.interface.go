package interfaces

import "context"

type ModifyCompetitorPoints interface {
	ModifyCompetitorPoints(ctx context.Context, categoryID, competitorID string, pointsNum int) error
}
