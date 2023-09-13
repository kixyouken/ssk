package routers

import (
	"ssk/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	admin := r.Group("/admin")
	{
		// 列表信息
		admin.GET("/table/:model", controllers.Get)
		// 详情信息
		admin.GET("/form/:model/:id", controllers.Read)
		// 添加信息
		admin.POST("/form/:model", controllers.Save)
		// 更新信息
		admin.PUT("/form/:model/:id", controllers.Update)
		// 删除信息
		admin.DELETE("/table/:model/:id", controllers.Delete)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
