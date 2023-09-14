package services

import (
	"encoding/json"
	"os"
	"ssk/jsons/forms"
	"ssk/jsons/models"
	"ssk/jsons/tables"
	"strings"

	"github.com/gin-gonic/gin"
)

type sFileService struct{}

var FileService = sFileService{}

// GetTableFile 获取 table 的 json 文件
//
//	@receiver s
//	@param c
//	@return *tables.BaseTable
func (s *sFileService) GetTableFile(c *gin.Context) *tables.BaseTable {
	path := c.Request.URL.Path
	pathList := strings.Split(strings.TrimLeft(path, "/"), "/")
	table := pathList[len(pathList)-1]
	tableFile := "./json/table/" + table + ".json"
	body, err := os.ReadFile(tableFile)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read JSON file"})
		return nil
	}
	tableJson := tables.BaseTable{}
	json.Unmarshal(body, &tableJson)

	return &tableJson
}

// GetTableOrders 获取 table 的排序信息
//
//	@receiver s
//	@param c
//	@param table
//	@return string
func (s *sFileService) GetTableOrders(c *gin.Context, table tables.BaseTable) string {
	orderList := []string{}
	for _, v := range table.Action.Orders {
		orderList = append(orderList, v.Field+" "+strings.ToUpper(v.Sort))
	}

	return strings.Join(orderList, ",")
}

// GetModelFile 获取 model 的 json 文件
//
//	@receiver s
//	@param c
//	@param model
//	@return *models.BaseModel
func (s *sFileService) GetModelFile(c *gin.Context, model string) *models.BaseModel {
	modelFile := "./json/model/" + model + ".json"
	body, err := os.ReadFile(modelFile)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read JSON file"})
		return nil
	}
	modelJson := models.BaseModel{}
	json.Unmarshal(body, &modelJson)

	return &modelJson
}

// GetModelColumns 获取 model 的字段列信息
//
//	@receiver s
//	@param c
//	@param model
//	@return []string
func (s *sFileService) GetModelColumns(c *gin.Context, model models.BaseModel) []string {
	column := []string{}
	for _, v := range model.Columns {
		if !strings.Contains(v.Field, ".") {
			column = append(column, "`"+model.Table.Name+"`.`"+v.Field+"`")
		} else {
			column = append(column, v.Field)
		}
	}

	return column
}

// GetModelJoins 获取 model 的关联信息
//
//	@receiver s
//	@param c
//	@param model
//	@return []string
func (s *sFileService) GetModelJoins(c *gin.Context, model models.BaseModel) []string {
	join := []string{}
	for _, v := range model.Table.Joins {
		join = append(join, strings.ToUpper(v.Join)+" JOIN "+v.Name+" ON "+v.Name+"."+v.Foreign+" = "+model.Table.Name+"."+v.Key)
	}

	return join
}

// GetFormFile 获取 form 的 json 文件
//
//	@receiver s
//	@param c
//	@return *forms.BaseForm
func (s *sFileService) GetFormFile(c *gin.Context) *forms.BaseForm {
	path := c.Request.URL.Path
	pathList := strings.Split(strings.TrimLeft(path, "/"), "/")
	form := pathList[len(pathList)-2]
	if pathList[0] == "admin" {
		form = pathList[len(pathList)-1]
	}
	formFile := "./json/form/" + form + ".json"
	body, err := os.ReadFile(formFile)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read JSON file"})
		return nil
	}
	formJson := forms.BaseForm{}
	json.Unmarshal(body, &formJson)

	return &formJson
}
