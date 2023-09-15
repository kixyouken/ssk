package services

import (
	"ssk/jsons/models"
	"strings"

	"github.com/gin-gonic/gin"
)

type sHandleService struct{}

var HandleService = sHandleService{}

// GetWiths 获取 with 子查询
//
//	@receiver s
//	@param c
//	@param result
//	@return []map
func (s *sHandleService) GetWiths(c *gin.Context, result []map[string]interface{}) []map[string]interface{} {
	table := FileService.GetTableFile(c)
	model := FileService.GetModelFile(c, table.Action.Bind.Model)
	for _, value := range result {
		for _, v := range model.Table.Withs {
			columns := s.GetWithsColumns(c, *model)
			order := s.GetWithsOrders(c, *model)

			withResult := []map[string]interface{}{}

			ModelService.GetAll(c, v.Name, &withResult, columns, order, map[string]interface{}{v.Foreign: value[v.Key]})
			value[v.Name] = withResult
		}
	}

	return result
}

// GetWithsColumns 获取 with 查询字段
//
//	@receiver s
//	@param c
//	@param model
//	@return []string
func (s *sHandleService) GetWithsColumns(c *gin.Context, model models.BaseModel) []string {
	columns := []string{}
	for _, value := range model.Table.Withs {
		for _, v := range value.Columns {
			columns = append(columns, v.Field)
		}
	}
	return columns
}

// GetWithsOrders 获取 with 排序信息
//
//	@receiver s
//	@param c
//	@param model
//	@return string
func (s *sHandleService) GetWithsOrders(c *gin.Context, model models.BaseModel) string {
	orders := []string{}
	for _, value := range model.Table.Withs {
		for _, v := range value.Orders {
			orders = append(orders, v.Field+" "+strings.ToUpper(v.Sort))
		}
	}

	return strings.Join(orders, ",")
}
