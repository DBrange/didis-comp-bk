package customerrors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Msg)
}

func ErrorResponse(err error, c *gin.Context) {
	var appErr AppError
	if errors.As(err, &appErr) {
		switch appErr.Code {

		case ErrCodeNotFound:
			c.JSON(http.StatusNotFound, appErr)

		case ErrCodeInsertionFailed:
			c.JSON(http.StatusConflict, appErr)

		case ErrCodeInvalidID:
			c.JSON(http.StatusBadRequest, appErr)

		case ErrCodeDuplicateKey:
			c.JSON(http.StatusConflict, appErr)

		case ErrCodeValidationFailed:
			c.JSON(http.StatusConflict, appErr)

		case ErrCodeSchemaViolation:
			c.JSON(http.StatusBadRequest, appErr)

		case ErrCodeUpdated:
			c.JSON(http.StatusBadRequest, appErr)

		case ErrCodeDeleted:
			c.JSON(http.StatusBadRequest, appErr)

		case ErrCodeGetJSON:
			c.JSON(http.StatusConflict, appErr)

		case ErrCodeTransaction:
			c.JSON(http.StatusServiceUnavailable, appErr)

		case ErrCodeStartSessionFailed:
			c.JSON(http.StatusServiceUnavailable, appErr)

		case ErrCodeAuthorizationHeader:
			c.JSON(http.StatusBadRequest, appErr)

		case ErrCodeHashedFailed:
			c.JSON(http.StatusBadRequest, appErr)

		case ErrCodeTokenSigned:
			c.JSON(http.StatusBadRequest, appErr)

		case ErrCodeSingedMethod:
			c.JSON(http.StatusBadRequest, appErr)

		case ErrCodeComparedHash:
			c.JSON(http.StatusBadRequest, appErr)

		case ErrCodeAuthorization:
			c.JSON(http.StatusUnauthorized, appErr)

		case ErrCodeAlreadyExits:
			c.JSON(http.StatusBadRequest, appErr)

		}

		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": ErrCodeNotFound, "error": fmt.Sprintf("unexpected error: %s", err.Error())})

}

type ErrorHandler func(error) AppError

func GenerateErrorHandler(code string, entityName string, msgTemplate string) ErrorHandler {
	return func(err error) AppError {
		return AppError{
			Code: code,
			Msg:  fmt.Sprintf(msgTemplate, entityName, err),
		}
	}
}

func CreateErrorHandlers(entityName string) map[error]ErrorHandler {
	return map[error]ErrorHandler{
		ErrNotFound:            GenerateErrorHandler(ErrCodeNotFound, entityName, "error when searching for %s: %v"),
		ErrInsertionFailed:     GenerateErrorHandler(ErrCodeInsertionFailed, entityName, "error inserting %s: %v"),
		ErrInvalidID:           GenerateErrorHandler(ErrCodeInvalidID, entityName, "invalid %s id format: %v"),
		ErrDuplicateKey:        GenerateErrorHandler(ErrCodeDuplicateKey, entityName, "error duplicate key for %s: %v"),
		ErrSchemaViolation:     GenerateErrorHandler(ErrCodeSchemaViolation, entityName, "error %s scheme type: %v"),
		ErrUpdated:             GenerateErrorHandler(ErrCodeUpdated, entityName, "error updating %s: %v"),
		ErrDeleted:             GenerateErrorHandler(ErrCodeDeleted, entityName, "error deleting %s: %v"),
		ErrValidationFailed:    GenerateErrorHandler(ErrCodeValidationFailed, entityName, "error validation %s: %v"),
		ErrGetJSON:             GenerateErrorHandler(ErrCodeGetJSON, entityName, "error binding json %s: %v"),
		ErrTransaction:         GenerateErrorHandler(ErrCodeTransaction, entityName, "error transaction %s: %v"),
		ErrStartSessionFailed:  GenerateErrorHandler(ErrCodeStartSessionFailed, entityName, "error start session %s: %v"),
		ErrAuthorizationHeader: GenerateErrorHandler(ErrCodeAuthorizationHeader, entityName, "error authorization %s: %v"),
		ErrHashedFailed:        GenerateErrorHandler(ErrCodeHashedFailed, entityName, "error hashing %s: %v"),
		ErrTokenSigned:         GenerateErrorHandler(ErrCodeTokenSigned, entityName, "error signing %s: %v"),
		ErrSingedMethod:        GenerateErrorHandler(ErrCodeSingedMethod, entityName, "error signing %s: %v"),
		ErrComparedHash:        GenerateErrorHandler(ErrCodeComparedHash, entityName, "error camparison %s: %v"),
		ErrAuthorization:       GenerateErrorHandler(ErrCodeAuthorization, entityName, "error authorization %s: %v"),
		ErrAlreadyExits:       GenerateErrorHandler(ErrCodeAlreadyExits, entityName, "already exits %s: %v"),
	}
}

func HandleError(err error, handlers map[error]ErrorHandler, msgTemplate string) error {
	for knownErr, handler := range handlers {
		if errors.Is(err, knownErr) {
			return handler(err)
		}
	}
	return fmt.Errorf("%s: %w", msgTemplate, err)
}

func HandleErrMsg(err error, name string, errMsg string) error {
	profileErrorHandlers := CreateErrorHandlers(name)
	return HandleError(err, profileErrorHandlers, errMsg)
}
