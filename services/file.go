package services

import (
	"encoding/json"
	"os"
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
	pathSlice := strings.Split(strings.TrimLeft(path, "/"), "/")
	table := pathSlice[len(pathSlice)-1]
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

// GetModelFileColumns 获取 model 的字段列信息
//
//	@receiver s
//	@param c
//	@param columns
//	@return []string
func (s *sFileService) GetModelFileColumns(c *gin.Context, columns []models.Columns) []string {
	column := []string{}
	for _, v := range columns {
		column = append(column, "`"+v.Name+"`")
	}

	return column
}
