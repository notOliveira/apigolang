package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/notOliveira/apigolang/schemas"
)

// @BasePath /api/v1

// @Summary Get opening
// @Description Get a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "query parameter 'id' is required")
		return
	}
	openings := []schemas.Opening{}
	if err := db.First(&openings, id).Error; err != nil {
		logger.Errorf("Error fetching Opening: %v", err.Error())
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}
	sendSuccess(ctx, "show-opening", openings)
}
