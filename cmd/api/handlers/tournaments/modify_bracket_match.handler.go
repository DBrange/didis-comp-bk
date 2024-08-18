package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	validate_util "github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ModifyBracketMatch(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	matchID := c.Param("matchID")

	competitors, err := modifyBracketMatchtBodyData(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.tournament.ModifyBracketMatch(ctx, matchID, competitors); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Match succsessfully updated!"})

}

func modifyBracketMatchtBodyData(c *gin.Context) ([]*dto.UpdateCompetitorMatchDTOReq, error) {
	var competitors struct {
		Competitors []*dto.UpdateCompetitorMatchDTOReq `json:"competitors"`
	}

	if err := c.ShouldBindJSON(&competitors); err != nil {
		err = fmt.Errorf("%w: error binding json: %v", customerrors.ErrGetJSON, err.Error())
		tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
		errMsgTemplate := "error getting tournament"
		return nil, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
	}

	// Validar la estructura excepto el campo Location
	err := validate_util.Validate.Struct(competitors)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
		errMsgTemplate := "error validation tournament"
		return nil, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
	}

	return competitors.Competitors, nil
}
