package services

import (
	"context"
	"fmt"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ProfileService) ModifyPassword(ctx context.Context, userID, newPassword, oldPassword string) error {
	password, err := s.profileQuerier.GetUserPasswordByID(ctx, userID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error getting profile password")
	}

	if !s.ComparePasswords(password, []byte(oldPassword)) {
		err := fmt.Errorf("error when searching for the user: %w", customerrors.ErrComparedHash)
		return customerrors.HandleErrMsg(err, "profile", "error passwords do not match")
	}

	passwordHashed, err := s.HashPassword(newPassword)
	if err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error when hashing password")
	}

	if err := s.profileQuerier.UpdateUserPassword(ctx, userID, passwordHashed); err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error updating password")
	}

	return nil
}
