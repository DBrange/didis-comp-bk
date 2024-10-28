package services

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/config"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/golang-jwt/jwt/v5"
)

func (s *ProfileService) RefreshToken(ctx context.Context, refreshToken string) (string, string, error) {
	userID, err := s.GetTokenUserID(refreshToken)
	if err != nil {
		return "", "", err
	}

	userDTO, err := s.profileQuerier.GetUserForRefreshToken(ctx, userID)
	if err != nil {
		return "", "", customerrors.HandleErrMsg(err, "profile", "error getting password")
	}

	roles := make([]string, len(userDTO.Roles))

	for i, roleID := range userDTO.Roles {
		roleStr, err := s.profileQuerier.GetRoleString(ctx, roleID)
		if err != nil {
			return "", "", customerrors.HandleErrMsg(err, "profile", "error getting role")
		}

		roles[i] = string(roleStr)
	}

	secret := []byte(config.Envs.JWTSecret)
	newToken, err := s.CreateJWT(secret, userDTO.ID, roles, config.Envs.JWTExpirationInSeconds)
	if err != nil {
		return "", "", customerrors.HandleErrMsg(err, "profile", "error when creating newToken")
	}

	newRefreshToken, err := s.CreateJWT(secret, userDTO.ID, roles, config.Envs.JWTRefreshExpirationInSeconds)
	if err != nil {
		return "", "", customerrors.HandleErrMsg(err, "profile", "error when creating newToken")
	}

	return newToken, newRefreshToken, nil
}

func (s *ProfileService) GetTokenUserID(refreshToken string) (string, error) {
	token, err := s.validateRefreshToken(refreshToken)
	if err != nil || !token.Valid {
		return "", s.permissionDenied("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["sub"].(string)
		if !ok {
			return "", s.permissionDenied("invalid user ID claim")
		}

		return userID, nil
	} else {
		return "", s.permissionDenied("invalid token claims")

	}

}

func (s *ProfileService) validateRefreshToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf("%w: unexpected signing method: %v", customerrors.ErrNotFound, token.Header["alg"])
			return nil, customerrors.HandleErrMsg(err, "profile", "error when creating newToken")
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func (s *ProfileService) permissionDenied(text string) error {
	err := fmt.Errorf("error authorization: %w", customerrors.ErrAuthorization)
	errMsgTemplate := fmt.Sprintf("error %s", text)
	return customerrors.HandleErrMsg(err, "auth", errMsgTemplate)
}
