package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ProfileService) GetUserPrimaryInfo(ctx context.Context, fromID, userToID string) (*dto.GetUserPrimatyInfoDTORes, error) {
	userDTO, err := s.profileQuerier.GetUserPrimaryData(ctx, userToID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "profile", "error getting primary info")
	}

	numberFollowers, err := s.profileQuerier.GetNumberFollowers(ctx, userToID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "profile", "error getting followers number")
	}

	isFollowing, err := s.profileQuerier.IsFollowing(ctx, fromID, userToID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "profile", "error getting isFollowing")
	}

	organizerID, err := s.profileQuerier.GetOrganizerIDByUserID(ctx, userToID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "profile", "error getting password")
	}

	primaryInfo := &dto.GetUserPrimatyInfoDTORes{
		User:        userDTO,
		Followers:   numberFollowers,
		IsFollowing: isFollowing,
		OrganizerID: organizerID,
	}

	return primaryInfo, nil
}
