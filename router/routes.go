package router

import (
	"github.com/LeonardoMedeiross/goJobs/controller"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/opening", controller.ShowOpeningHandler)

		v1.POST("/opening", controller.CreateOpeningHandler)

		v1.DELETE("/opening", controller.DeleteOpeningHandler)

		v1.PUT("/opening", controller.UpdateOpeningHandler)

		v1.GET("/openings", controller.ListOpeningsHandler)
	}
}
