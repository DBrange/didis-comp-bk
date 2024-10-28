package services

import (
	"context"
	"time"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ProfileService) GetUserFollowers(ctx context.Context, userID string,name string, limit int, lastCreatedAt *time.Time) (*dto.GetUserFollowersDTORes, error) {
	userFollowers, err := s.profileQuerier.GetUserFollowers(ctx, userID,name, limit, lastCreatedAt)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "profile", "error getting followers")
	}

	return userFollowers, nil
}
