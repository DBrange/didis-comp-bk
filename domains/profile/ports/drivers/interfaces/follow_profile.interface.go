package interfaces

import "context"

type FollowProfile interface {
	FollowProfile(ctx context.Context, fromUserID, toUserID string) error
}
