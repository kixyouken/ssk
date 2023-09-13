package routers

import (
	"ssk/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	admin := r.Group("/admin")
	{
		admin.GET("/table/:model", controllers.Get)
		admin.GET("/form/:model/:id", controllers.Read)
		admin.POST("/form/:model", controllers.Save)
		admin.PUT("/form/:model/:id", controllers.Update)
		admin.DELETE("/table/:model/:id", controllers.Delete)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
