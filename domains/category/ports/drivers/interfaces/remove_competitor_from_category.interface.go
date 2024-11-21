package interfaces

import "context"

type RemoveCompetitorFromCategory interface {
	RemoveCompetitorFromCategory(ctx context.Context, categoryID,competitorID string) error
}
