package main

import (
	"github.com/notOliveira/apigolang/config"
	"github.com/notOliveira/apigolang/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")

	err := config.Init()

	if err == nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize the router
	router.Initialize()
}