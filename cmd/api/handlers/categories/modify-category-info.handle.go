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

func (h *Handler) ModifyCategoryInfo(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	categoryID := c.Param("categoryID")

	categoryInfoDTO, err := modifyCategoryInfoValidateQueries(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.category.ModifyCategoryInfo(ctx, categoryID, categoryInfoDTO); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Category successfully updated!"})
}

func modifyCategoryInfoValidateQueries(c *gin.Context) (*dto.UpdateCategoryDTOReq, error) {
	var categoryDTO dto.UpdateCategoryDTOReq
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
