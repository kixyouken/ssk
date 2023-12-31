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
	tableFile := "./json/table/" + pathList[2] + ".json"
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

// GetTableDistincts 获取去重字段
//
//	@receiver s
//	@param c
//	@param table
//	@return []string
func (s *sFileService) GetTableDistincts(c *gin.Context, table tables.BaseTable) []string {
	distinctList := []string{}
	for _, v := range table.Action.Bind.Filter.Distinct {
		distinctList = append(distinctList, v.Field)
	}

	return distinctList
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
	if model.Columns == nil || len(model.Columns) == 0 {
		column = append(column, model.Table.Name+".*")
	} else {
		for _, v := range model.Columns {
			if !strings.Contains(v.Field, ".") {
				column = append(column, "`"+model.Table.Name+"`.`"+v.Field+"`")
			} else {
				column = append(column, v.Field)
			}
		}
	}

	joinColumn := ModelService.JoinColumns(c, model)
	column = append(column, joinColumn...)
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
	for _, value := range model.Table.Joins {
		joinTable := strings.ToUpper(value.Join) + " JOIN " + value.Name + " ON " + value.Name + "." + value.Foreign + " = " + model.Table.Name + "." + value.Key
		if value.Wheres != nil {
			joinWhere := []string{}
			for _, v := range value.Wheres {
				joinWhere = append(joinWhere, value.Name+"."+v.Field+s.WhereType(v.Search)+"'"+v.Value+"'")
			}
			if joinWhere != nil {
				joinTable += " AND ( " + strings.Join(joinWhere, " AND ") + " )"
			}
		}
		join = append(join, joinTable)
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
	formFile := "./json/form/" + pathList[2] + ".json"
	body, err := os.ReadFile(formFile)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read JSON file"})
		return nil
	}
	formJson := forms.BaseForm{}
	json.Unmarshal(body, &formJson)

	return &formJson
}

func (s *sFileService) WhereType(where string) string {
	switch where {
	case "eq":
		where = " = "
	case "neq":
		where = " <> "
	case "lt":
		where = " < "
	case "elt":
		where = " <= "
	case "gt":
		where = " > "
	case "egt":
		where = " >= "
	}
	return where
}
