package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	
	// Initialize Gin router
	r := gin.Default()

	InitializeRoutes(r)

	// Run the server
	r.Run(":8080")
	if err := r.Run(); err != nil {
		panic(err)
	}
}