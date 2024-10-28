package interfaces

import "context"

type GetNumberFollowers interface {
	GetNumberFollowers(ctx context.Context, userID string) (int, error)
}
