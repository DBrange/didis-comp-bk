package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddCompetitorInCategory(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	categoryID := c.Param("categoryID")
	competitorID := c.Param("competitorID")

	if err := h.category.AddCompetitorInCategory(ctx, categoryID, competitorID); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Competitor successfully added!"})
}

// func organizeCategoryBodyData(c *gin.Context) (*dto.OrganizeCategoryDTOReq, error) {
// 	var categoryInfoDTO dto.OrganizeCategoryDTOReq
// 	if err := c.ShouldBindJSON(&categoryInfoDTO); err != nil {
// 		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrGetJSON, err.Error())
// 		categoryErrorHandlers := customerrors.CreateErrorHandlers("category")
// 		errMsgTemplate := "error getting category"
// 		return nil, customerrors.HandleError(err, categoryErrorHandlers, errMsgTemplate)
// 	}

// 	err := utils.Validate.Struct(categoryInfoDTO)
// 	if err != nil {
// 		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
// 		categoryErrorHandlers := customerrors.CreateErrorHandlers("category")
// 		errMsgTemplate := "error validation category"
// 		return nil, customerrors.HandleError(err, categoryErrorHandlers, errMsgTemplate)
// 	}

// 	return &categoryInfoDTO, nil
// }
