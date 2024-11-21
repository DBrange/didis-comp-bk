package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	loginDTO, err := loginDataBody(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	user, token, refreshToken, err := h.profile.Login(ctx, loginDTO)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	type Data struct {
		ID           string                           `json:"id"`
		FirstName    string                           `json:"first_name"`
		LastName     string                           `json:"last_name"`
		Username     string                           `json:"username"`
		Image        string                           `json:"image"`
		Roles        []string                         `json:"roles"`
		Sports       []models.SPORT                   `json:"sports"`
		Organizer    *dto.GetUserForLoginOrganizerDTO `json:"organizer"`
		Token        string                           `json:"token"`
		RefreshToken string                           `json:"refresh_token"`
	}

	data := Data{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		Image:        user.Image,
		Roles:        user.Roles,
		Sports:       user.Sports,
		Organizer:    user.Organizer,
		Token:        token,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusCreated, gin.H{"data": data, "status": http.StatusOK, "message": "You have successfully logged in!"})

}

func loginDataBody(c *gin.Context) (*dto.LoginDTOReq, error) {
	var loginDTO dto.LoginDTOReq
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		err = fmt.Errorf("%w: error binding json: %v", customerrors.ErrGetJSON, err.Error())
		profileErrorHandlers := customerrors.CreateErrorHandlers("profile")
		errMsgTemplate := "error getting profile"
		return nil, customerrors.HandleError(err, profileErrorHandlers, errMsgTemplate)
	}

	err := utils.Validate.Struct(loginDTO)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		profileErrorHandlers := customerrors.CreateErrorHandlers("profile")
		errMsgTemplate := "error validation profile"
		return nil, customerrors.HandleError(err, profileErrorHandlers, errMsgTemplate)

	}

	return &loginDTO, nil
}
