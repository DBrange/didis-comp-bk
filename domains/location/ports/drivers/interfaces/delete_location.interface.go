package interfaces

import "context"

type DeleteLocation interface {
	DeleteLocation(ctx context.Context, locationID string) error
}
