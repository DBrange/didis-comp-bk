package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ModifyPassword(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	userID := c.Param("userID")

	newPassword, oldPassword, err := modifyPasswordBody(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.profile.ModifyPassword(ctx, userID, newPassword, oldPassword); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "The password has been successfully modified!"})

}

func modifyPasswordBody(c *gin.Context) (string, string, error) {
	type NewPassword struct {
		OldPassword string `json:"old_password" validate:"password"`
		NewPassword string `json:"new_password" validate:"password"`
	}

	var Passwords NewPassword

	if err := c.ShouldBindJSON(&Passwords); err != nil {
		err = fmt.Errorf("%w: error getting the json: %v", customerrors.ErrGetJSON, err.Error())
		if errors.Is(err, customerrors.ErrValidationFailed) {
			appErr := customerrors.AppError{
				Code: customerrors.ErrCodeGetJSON,
				Msg:  fmt.Sprintf("error binding: %v", err),
			}
			return "", "", appErr
		}
		return "", "", fmt.Errorf("error getting the json: %w", err)
	}

	err := utils.Validate.Struct(Passwords)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		if errors.Is(err, customerrors.ErrValidationFailed) {
			appErr := customerrors.AppError{
				Code: customerrors.ErrCodeValidationFailed,
				Msg:  fmt.Sprintf("error validation: %v", err),
			}
			return "", "", appErr
		}
		return "", "", fmt.Errorf("error validation: %w", err)
	}

	return Passwords.NewPassword, Passwords.OldPassword, nil

}
