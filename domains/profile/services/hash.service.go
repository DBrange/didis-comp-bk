package services

import (
	"fmt"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *ProfileService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		err := fmt.Errorf("%w: error when hasing password: %s", customerrors.ErrHashedFailed, err.Error())
		return "", customerrors.HandleErrMsg(err, "profile", "error hashing password")
	}

	return string(hash), nil
}

func (s *ProfileService) ComparePasswords(hashed string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)
	return err == nil
}
