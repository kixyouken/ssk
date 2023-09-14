package controllers

import (
	"ssk/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	table := services.FileService.GetTableFile(c)
	model := services.FileService.GetModelFile(c, table.Action.Bind.Model)
	column := services.FileService.GetModelColumns(c, *model)
	join := services.FileService.GetModelJoins(c, *model)
	count := services.ModelService.GetCount(c, model.Table.Name, join...)

	result := []map[string]interface{}{}
	if count > 0 {
		services.ModelService.GetPage(c, model.Table.Name, &result, column, model.Table.Name+".id DESC", join...)
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    result,
		"count":   count,
	})
}

func Read(c *gin.Context) {
	form := services.FileService.GetFormFile(c)
	model := services.FileService.GetModelFile(c, form.Action.Bind.Model)
	column := services.FileService.GetModelColumns(c, *model)

	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)
	result := map[string]interface{}{}
	services.ModelService.GetID(c, model.Table.Name, idInt, &result, column)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    result,
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
