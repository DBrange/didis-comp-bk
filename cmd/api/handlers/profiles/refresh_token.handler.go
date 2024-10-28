package handlers

import (
	"context"
	"fmt"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RefreshToken(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	refreshToken, err := refreshTokenDataBody(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	token, refreshToken, err := h.profile.RefreshToken(ctx, refreshToken)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	type Data struct {
		Token        string   `json:"token"`
		RefreshToken string   `json:"refresh_token"`
	}

	data := Data{
		Token:        token,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusCreated, gin.H{"data": data, "status": http.StatusOK, "message": "You have successfully logged in!"})

}

func refreshTokenDataBody(c *gin.Context) (string, error) {
var refreshTokenDTO struct {
    RefreshToken string `json:"refresh_token"`
}

	if err := c.ShouldBindJSON(&refreshTokenDTO); err != nil {
		err = fmt.Errorf("%w: error binding json: %v", customerrors.ErrGetJSON, err.Error())
		profileErrorHandlers := customerrors.CreateErrorHandlers("profile")
		errMsgTemplate := "error getting profile"
		return "", customerrors.HandleError(err, profileErrorHandlers, errMsgTemplate)
	}

	err := utils.Validate.Struct(refreshTokenDTO)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		profileErrorHandlers := customerrors.CreateErrorHandlers("profile")
		errMsgTemplate := "error validation profile"
		return "", customerrors.HandleError(err, profileErrorHandlers, errMsgTemplate)

	}

	return refreshTokenDTO.RefreshToken, nil
}
