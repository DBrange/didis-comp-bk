package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/config"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ProfileService) Login(ctx context.Context, loginDTO *dto.LoginDTOReq) (*dto.GetUserForLoginDTO, string, string, error) {
	userDTO, err := s.profileQuerier.GetUserForLogin(ctx, loginDTO.Username)
	if err != nil {
		return nil, "", "", customerrors.HandleErrMsg(err, "profile", "error getting password")
	}

	organizerID, err := s.profileQuerier.GetOrganizerIDByUserID(ctx, userDTO.ID)
	if err != nil {
		return nil, "", "", customerrors.HandleErrMsg(err, "profile", "error getting password")
	}

	sports, err := s.profileQuerier.GetUserAllCompetitorSports(ctx, userDTO.ID)
	if err != nil {
		return nil, "", "", customerrors.HandleErrMsg(err, "profile", "error getting sports")
	}

	// Add organizerID if != nil
	userDTO.OrganizerID = organizerID

	if !s.ComparePasswords(userDTO.Password, []byte(loginDTO.Password)) {
		return nil, "", "", customerrors.HandleErrMsg(err, "profile", "error passwords do not match")
	}

	// roles, err := s.profileQuerier.GetUserRoles(ctx, userID)
	// if err != nil {
	// 	return nil,  "","",customerrors.HandleErrMsg(err, "profile", "error getting profile roles")
	// }

	roles := make([]string, len(userDTO.Roles))

	for i, roleID := range userDTO.Roles {
		roleStr, err := s.profileQuerier.GetRoleString(ctx, roleID)
		if err != nil {
			return nil, "", "", customerrors.HandleErrMsg(err, "profile", "error getting role")
		}

		roles[i] = string(roleStr)
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := s.CreateJWT(secret, userDTO.ID, roles, config.Envs.JWTExpirationInSeconds)
	if err != nil {
		return nil, "", "", customerrors.HandleErrMsg(err, "profile", "error when creating token")
	}

	refreshToken, err := s.CreateJWT(secret, userDTO.ID, roles, config.Envs.JWTRefreshExpirationInSeconds)
	if err != nil {
		return nil, "", "", customerrors.HandleErrMsg(err, "profile", "error when creating token")
	}

	userDTO.Roles = roles
	userDTO.Sports = sports

	return userDTO, token, refreshToken, nil
}
