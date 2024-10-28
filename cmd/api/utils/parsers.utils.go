package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

func ParseToFloat(c *gin.Context, query string) (float64, error) {
	queryStr := c.Query(query)

	queryFloat, err := strconv.ParseFloat(queryStr, 64)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		ErrorHandlers := customerrors.CreateErrorHandlers("")
		errMsgTemplate := "error validation "
		return 0, customerrors.HandleError(err, ErrorHandlers, errMsgTemplate)
	}
	return queryFloat, nil
}

func ParseToTime(c *gin.Context, query string) (*time.Time, error) {
	// Obtener el valor de la query string
	queryStr := c.Query(query)

	if queryStr == "" {
		return nil, nil
	}

	// Reemplazar el espacio entre la fecha y la hora con 'T'
	queryStr = strings.ReplaceAll(queryStr, " ", "T")

	// Intentar parsear la fecha/hora según el formato dado
	parsedTime, err := time.Parse(time.RFC3339Nano, queryStr) // Usa el formato RFC3339
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		ErrorHandlers := customerrors.CreateErrorHandlers("")
		errMsgTemplate := "error validation"
		return nil, customerrors.HandleError(err, ErrorHandlers, errMsgTemplate)
	}

	return &parsedTime, nil
}
