package routers

import (
	"ssk/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	admin := r.Group("/admin")
	{
		// 表格配置
		admin.GET("/table/:setting", controllers.GetTable)
		// 表单配置
		admin.GET("/form/:setting", controllers.GetForm)
		// 模型配置
		admin.GET("/model/:setting", controllers.GetModel)
	}

	api := r.Group("/api")
	{
		// 列表信息
		api.GET("/table/:model", controllers.Get)
		// 详情信息
		api.GET("/form/:model/:id", controllers.Read)
		// 添加信息
		api.POST("/form/:model", controllers.Save)
		// 更新信息
		api.PUT("/form/:model/:id", controllers.Update)
		// 删除信息
		api.DELETE("/table/:model/:id", controllers.Delete)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
