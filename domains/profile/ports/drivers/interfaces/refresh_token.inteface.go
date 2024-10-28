package interfaces

import (
	"context"
)

type RefreshToken interface {
	RefreshToken(ctx context.Context, refreshToken string) (string, string, error)
}
