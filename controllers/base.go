package controllers

import (
	"ssk/services"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	table := services.FileService.GetTableFile(c)
	model := services.FileService.GetModelFile(c, table.Action.Bind.Model)
	join := services.FileService.GetModelJoins(c, *model)
	column := services.FileService.GetModelColumns(c, *model)
	order := services.FileService.GetTableOrders(c, *table)

	where := ""
	if model.Table.Wheres != nil {
		where = services.HandleService.GetModelWheres(c, *model)
	}

	var count int64
	result := []map[string]interface{}{}
	if table.Action.Bind.Filter.Distinct != nil {
		fieldList := services.FileService.GetTableDistincts(c, *table)
		fields := strings.Join(fieldList, ",")
		count = services.ModelService.GetDistinctCount(c, model.Table.Name, fields)
		services.ModelService.GetDistinct(c, model.Table.Name, &result, fields)
	} else {
		count = services.ModelService.GetCount(c, model.Table.Name, join, where)
		if table.Action.Count > 0 && int64(table.Action.Count) < count {
			count = int64(table.Action.Count)
		}

		if count > 0 {
			services.ModelService.GetPage(c, model.Table.Name, &result, column, order, join, where)
		}

		if model.Table.WithsCount != nil {
			result = services.HandleService.GetWithsCount(c, result, *model)
		}

		if model.Table.Withs != nil {
			result = services.HandleService.GetWiths(c, result, *model)
		}
	}

	if model.Columns != nil {
		services.HandleService.GetFieldText(c, *model, result)
		services.HandleService.GetFieldFormat(c, *model, result)
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
	join := services.FileService.GetModelJoins(c, *model)

	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)
	result := []map[string]interface{}{}
	services.ModelService.GetID(c, idInt, model.Table.Name, &result, column, join)

	if model.Table.WithsCount != nil {
		services.HandleService.GetWithsCount(c, result, *model)
	}

	if model.Table.Withs != nil {
		services.HandleService.GetWiths(c, result, *model)
	}

	if model.Columns != nil {
		services.HandleService.GetFieldText(c, *model, result)
		services.HandleService.GetFieldFormat(c, *model, result)
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    result[0],
	})
}

func Save(c *gin.Context) {
	form := services.FileService.GetFormFile(c)
	model := services.FileService.GetModelFile(c, form.Action.Bind.Model)
	param := map[string]interface{}{}
	c.ShouldBind(&param)
	services.ModelService.SetCreate(c, model.Table.Name, param)

	// TODO: 如何返回 id
	c.JSON(200, gin.H{
		"message": "Save",
		"param":   param,
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
