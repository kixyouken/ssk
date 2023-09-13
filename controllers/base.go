package controllers

import (
	"ssk/services"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	table := services.FileService.GetTableFile(c)
	model := services.FileService.GetModelFile(c, table.Action.Bind.Table)
	column := services.FileService.GetModelFileColumns(c, model.Columns)
	result := []map[string]interface{}{}
	services.ModelService.GetPage(c, model.Table.Name, &result, column, "")
	count := services.ModelService.GetCount(c, model.Table.Name)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    result,
		"count":   count,
		"table":   table,
		"model":   model,
	})
}

func Read(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Read",
	})
}

func Save(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Save",
	})
}

func Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update",
	})
}

func Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete",
	})
}
