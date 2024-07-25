package utils

import (
	"fmt"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func ParseErr(name string) error {
	err := fmt.Errorf("%w: error parsing %s", customerrors.ErrValidationFailed, name)
	profileErrorHandlers := customerrors.CreateErrorHandlers("profile")
	errMsgTemplate := "error parsing " + name
	return customerrors.HandleError(err, profileErrorHandlers, errMsgTemplate)
}
