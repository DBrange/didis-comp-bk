package services

import (
	"fmt"
	"time"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/golang-jwt/jwt/v5"
)

func (s *ProfileService) CreateJWT(secret []byte, userID string, roles []string, expirationSec int64) (string, error) {
	expiration := time.Second * time.Duration(expirationSec)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userID,
		"roles": roles,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		err := fmt.Errorf("%w: error when signed token: %s", customerrors.ErrTokenSigned, err.Error())
		return "", customerrors.HandleErrMsg(err, "profile", "error signed token")
	}

	return tokenString, nil
}
