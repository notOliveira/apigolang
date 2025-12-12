package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/notOliveira/apigolang/schemas"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "query parameter 'id' is required")
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("Error fetching Opening: %v", err.Error())
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	// Update fields
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	// Save updated opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("Error updating Opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
	sendSuccess(ctx, "update-opening", opening)
}
