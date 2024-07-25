package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) AddCompetitorInLeague(c *gin.Context) {
	// ctx, cancel := context.WithCancel(c.Request.Context())
	// defer cancel()

	// tournamentID := c.Param("tournamentID")
	// leagueID := c.Param("leagueID")

	// if err := h.league.AddTournamentInLeague(ctx, leagueID, tournamentID); err != nil {
	// 	customerrors.ErrorResponse(err, c)
	// 	return
	// }

	// c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Competitor successfully added!"})
}

// func organizeLeagueBodyData(c *gin.Context) (*dto.OrganizeLeagueDTOReq, error) {
// 	var leagueInfoDTO dto.OrganizeLeagueDTOReq
// 	if err := c.ShouldBindJSON(&leagueInfoDTO); err != nil {
// 		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrGetJSON, err.Error())
// 		leagueErrorHandlers := customerrors.CreateErrorHandlers("league")
// 		errMsgTemplate := "error getting league"
// 		return nil, customerrors.HandleError(err, leagueErrorHandlers, errMsgTemplate)
// 	}

// 	err := utils.Validate.Struct(leagueInfoDTO)
// 	if err != nil {
// 		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
// 		leagueErrorHandlers := customerrors.CreateErrorHandlers("league")
// 		errMsgTemplate := "error validation league"
// 		return nil, customerrors.HandleError(err, leagueErrorHandlers, errMsgTemplate)
// 	}

// 	return &leagueInfoDTO, nil
// }
