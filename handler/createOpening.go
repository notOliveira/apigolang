package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating CreateOpeningRequest: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "Invalid request: "+err.Error())
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("Error creating Opening: %v", err.Error())
		ctx.JSON(500, gin.H{"error": "Error creating opening"})
		return
	}

}