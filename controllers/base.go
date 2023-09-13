package controllers

import (
	"ssk/models"
	"ssk/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	path := c.Request.URL.Path
	pathSlice := strings.Split(strings.TrimLeft(path, "/"), "/")
	model := []models.BaseModel{}

	services.GetPage(c, pathSlice[1], &model, "*", "")
	c.JSON(200, gin.H{
		"message": "Get",
		"data":    model,
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
