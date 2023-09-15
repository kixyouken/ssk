package controllers

import (
	"ssk/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	table := services.FileService.GetTableFile(c)
	count := services.ModelService.GetCount(c)
	if table.Action.Count > 0 && int64(table.Action.Count) < count {
		count = int64(table.Action.Count)
	}

	result := []map[string]interface{}{}
	if count > 0 {
		services.ModelService.GetPage(c, &result)
	}

	model := services.FileService.GetModelFile(c, table.Action.Bind.Model)
	if model.Table.Withs != nil {
		result = services.HandleService.GetWiths(c, result)
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    result,
		"count":   count,
	})
}

func Read(c *gin.Context) {
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)
	result := map[string]interface{}{}
	services.ModelService.GetID(c, &result, idInt)

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
