package interfaces

import "context"

type ModifyPassword interface {
	ModifyPassword(ctx context.Context, userID, newPassword, oldPassword string) error
}
