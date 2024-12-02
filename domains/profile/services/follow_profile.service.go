package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ProfileService) FollowProfile(ctx context.Context, fromUserID, toUserID string) error {
	followerDTO := &dto.CreateFollowerDTOReq{
		From:   fromUserID,
		ToUser: &toUserID,
	}

	if err := s.profileQuerier.FollowOrUnfollow(ctx, followerDTO); err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error you were already following him")
	}

	return nil
}
