package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/config"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ProfileService) Login(ctx context.Context, loginDTO *dto.LoginDTOReq) (string, string, error) {
	password, userID, err := s.profileQueryer.GetUserPasswordForLogin(ctx, loginDTO.Username)
	if err != nil {
		return "", "", customerrors.HandleErrMsg(err, "profile", "error getting password")
	}

	if !s.ComparePasswords(password, []byte(loginDTO.Password)) {
		return "", "", customerrors.HandleErrMsg(err, "profile", "error passwords do not match")
	}

	roles, err := s.profileQueryer.GetUserRoles(ctx, userID)
	if err != nil {
		return "", "", customerrors.HandleErrMsg(err, "profile", "error getting profile roles")
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := s.CreateJWT(secret, userID, roles, config.Envs.JWTExpirationInSeconds)
	if err != nil {
		return "", "", customerrors.HandleErrMsg(err, "profile", "error when creating token")
	}

	refreshToken, err := s.CreateJWT(secret, userID, roles, config.Envs.JWTRefreshExpirationInSeconds)
	if err != nil {
		return "", "", customerrors.HandleErrMsg(err, "profile", "error when creating token")
	}

	return token, refreshToken, nil
}
