package services

import (
	"ssk/jsons/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type sHandleService struct{}

var HandleService = sHandleService{}

// GetWithsCount 获取 withCount 子查询
//
//	@receiver s
//	@param c
//	@param result
//	@return []map
func (s *sHandleService) GetWithsCount(c *gin.Context, result []map[string]interface{}, model models.BaseModel) []map[string]interface{} {
	for _, value := range result {
		for _, v := range model.Table.WithsCount {
			withCount := ModelService.GetWithsCount(c, v.Name, map[string]interface{}{v.Foreign: value[v.Key]})
			value[v.Name+"_count"] = withCount
		}
	}

	return result
}

// GetWiths 获取 with 子查询
//
//	@receiver s
//	@param c
//	@param result
//	@return []map
func (s *sHandleService) GetWiths(c *gin.Context, result []map[string]interface{}, model models.BaseModel) []map[string]interface{} {
	for _, value := range result {
		for _, v := range model.Table.Withs {
			columns := s.GetWithsColumns(c, model)
			order := s.GetWithsOrders(c, model)

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

// GetFieldText 转文字
//
//	@receiver s
//	@param c
//	@param model
//	@param result
//	@return []map
func (s *sHandleService) GetFieldText(c *gin.Context, model models.BaseModel, result []map[string]interface{}) []map[string]interface{} {
	for _, column := range model.Columns {
		if column.Attrs != nil {
			for _, attr := range column.Attrs {
				for _, value := range result {
					in, _ := strconv.Atoi(attr.In)
					if value[column.Field] == int32(in) {
						value[column.Field+"_text"] = attr.Out
					}
				}
			}
		}
	}
	return result
}

// GetFieldFormat 格式化时间
//
//	@receiver s
//	@param c
//	@param model
//	@param result
//	@return []map
func (s *sHandleService) GetFieldFormat(c *gin.Context, model models.BaseModel, result []map[string]interface{}) []map[string]interface{} {
	for _, column := range model.Columns {
		if column.Format != "" {
			format := ""
			switch column.Format {
			case "Y-m-d":
				format = "2006-01-02"
			case "Y-m-d H":
				format = "2006-01-02 15"
			case "Y-m-d H:i":
				format = "2006-01-02 15:04"
			case "Y-m-d H:i:s":
				format = "2006-01-02 15:04:05"
			}
			for _, value := range result {
				date, _ := value[column.Field].(time.Time)
				value[column.Field] = date.Format(format)
			}
		}
	}

	return result
}

func (s *sHandleService) GetModelWheres(c *gin.Context, model models.BaseModel) string {
	where := ""
	for _, v := range model.Table.Wheres {
		if !strings.Contains(v.Field, ".") {
			where = model.Table.Name + "." + v.Field + " " + v.Search + " '" + v.Value + "'"
		} else {
			where = v.Field + " " + v.Search + " '" + v.Value + "'"
		}
	}

	return where
}
