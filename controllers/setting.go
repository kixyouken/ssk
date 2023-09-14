package controllers

import (
	"ssk/services"

	"github.com/gin-gonic/gin"
)

func GetTable(c *gin.Context) {
	table := services.FileService.GetTableFile(c)
	if table != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    table,
	})
}

func GetForm(c *gin.Context) {
	form := services.FileService.GetFormFile(c)
	if form != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    form,
	})
}

func GetModel(c *gin.Context) {
	table := services.FileService.GetTableFile(c)
	model := services.FileService.GetModelFile(c, table.Action.Bind.Model)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    model,
	})
}
