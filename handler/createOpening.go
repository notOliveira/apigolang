package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/notOliveira/apigolang/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating CreateOpeningRequest: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "Invalid request: "+err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Salary:   request.Salary,
		Link:     request.Link,
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("Error creating Opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)

}
