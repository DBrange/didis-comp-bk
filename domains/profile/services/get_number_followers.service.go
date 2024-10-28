package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ProfileService) GetNumberFollowers(ctx context.Context, userID string) (int, error) {
	numberFollowers, err := s.profileQuerier.GetNumberFollowers(ctx, userID)
	if err != nil {
		return 0, customerrors.HandleErrMsg(err, "profile", "error when get number of followersc")
	}

	return numberFollowers, nil
}
