package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/notOliveira/apigolang/schemas"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "Query parameter 'id' is required")
		return
	}
	openings := []schemas.Opening{}
	if err := db.First(&openings, id).Error; err != nil {
		logger.Errorf("Error fetching Opening: %v", err.Error())
		sendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}
	sendSuccess(ctx, "show-opening", openings)
}