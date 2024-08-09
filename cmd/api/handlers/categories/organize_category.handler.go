package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) OrganizeCategory(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	organizerID := c.Param("organizerID")

	categoryInfoDTO, err := organizeCategoryBodyData(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.category.OrganizeCategory(ctx, organizerID, categoryInfoDTO); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "New category successfully created !"})
}

func organizeCategoryBodyData(c *gin.Context) (*dto.CreateCategoryDTOReq, error) {
	var categoryDTO dto.CreateCategoryDTOReq
	if err := c.ShouldBindJSON(&categoryDTO); err != nil {
		err = fmt.Errorf("%w: error binding json: %v", customerrors.ErrGetJSON, err.Error())
		categoryErrorHandlers := customerrors.CreateErrorHandlers("category")
		errMsgTemplate := "error getting category"
		return nil, customerrors.HandleError(err, categoryErrorHandlers, errMsgTemplate)
	}

	err := utils.Validate.Struct(categoryDTO)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		categoryErrorHandlers := customerrors.CreateErrorHandlers("category")
		errMsgTemplate := "error validation category"
		return nil, customerrors.HandleError(err, categoryErrorHandlers, errMsgTemplate)
	}

	return &categoryDTO, nil
}
