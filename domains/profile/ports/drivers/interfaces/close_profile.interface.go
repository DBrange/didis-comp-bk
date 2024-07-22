package interfaces

import "context"

type CloseProfile interface {
	CloseProfile(ctx context.Context, userID string) error
}
