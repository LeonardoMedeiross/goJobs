package main

import (
	"github.com/LeonardoMedeiross/goJobs/database"
	"github.com/LeonardoMedeiross/goJobs/router"
)

var logger database.Logger

func main() {
	logger = *database.GetLogger("main")

	//initialize config
	err := database.Init()
	if err != nil {
		logger.Errorf("config initialization error : %v", err)
		return
	}
	//initialize router
	router.Initialize()

}
