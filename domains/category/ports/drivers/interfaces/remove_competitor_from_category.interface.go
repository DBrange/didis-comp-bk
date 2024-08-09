package interfaces

import "context"

type RemoveCompetitorFromCategory interface {
	RemoveCompetitorFromCategory(ctx context.Context, categoryRegistrationID string) error
}
