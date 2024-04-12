package main

import (
	"github.com/reishenrique/job-gopportunities/config"
	"github.com/reishenrique/job-gopportunities/router"
) 

var (
	logger *config.Logger
)

func main() {	
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.InitDatabase()
	if err != nil {
		logger.Errorf("Initialization Database Error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}