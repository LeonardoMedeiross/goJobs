package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	//initialize router
	r := gin.Default()

	//initialize routes
	InitializeRoutes(r)

	//run the server
	r.Run() // listen and serve on 0.0.0.0:8080
}
