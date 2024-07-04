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

		}

		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"status": ErrCodeNotFound, "error": "something has gone wrong"})

}
