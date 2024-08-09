package utils

import (
	"fmt"
	"strconv"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func ParseToBool(c *gin.Context, query string) (bool, error) {
	queryStr := c.Query(query)

	queryBol, err := strconv.ParseBool(queryStr)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		ErrorHandlers := customerrors.CreateErrorHandlers("")
		errMsgTemplate := "error validation "
		return false, customerrors.HandleError(err, ErrorHandlers, errMsgTemplate)
	}
	return queryBol, nil
}

func ParseToInt(c *gin.Context, query string) (int, error) {
	queryStr := c.Query(query)

	queryInt, err := strconv.Atoi(queryStr)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		ErrorHandlers := customerrors.CreateErrorHandlers("")
		errMsgTemplate := "error validation "
		return 0, customerrors.HandleError(err, ErrorHandlers, errMsgTemplate)
	}
	return queryInt, nil
}
