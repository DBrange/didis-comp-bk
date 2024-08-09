package interfaces

import "context"

type AddCompetitorInCategory interface {
	AddCompetitorInCategory(ctx context.Context, categoryID, competitorID string) error
}