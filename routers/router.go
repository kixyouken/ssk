package routers

import (
	"ssk/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/:controller", controllers.Get)
		api.GET("/:controller/:id", controllers.Read)
		api.POST("/:controller", controllers.Save)
		api.PUT("/:controller/:id", controllers.Update)
		api.DELETE("/:controller/:id", controllers.Delete)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
