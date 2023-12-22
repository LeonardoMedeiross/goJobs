package router

import "github.com/gin-gonic/gin"

func Initialize() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Leonardo Vai ser Dev na gringa com salario de 100k de dolls por ano "})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
