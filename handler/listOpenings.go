package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/notOliveira/apigolang/schemas"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error fetching openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}
