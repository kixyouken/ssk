package services

import (
	"ssk/jsons/models"
	"ssk/jsons/tables"
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
			if v.Type == "hasOne" {
				value[v.Name] = withResult[0]
			} else {
				value[v.Name] = withResult
			}
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
			column.Format = strings.Replace(column.Format, "Y", "2006", -1)
			column.Format = strings.Replace(column.Format, "m", "01", -1)
			column.Format = strings.Replace(column.Format, "d", "02", -1)
			column.Format = strings.Replace(column.Format, "H", "15", -1)
			column.Format = strings.Replace(column.Format, "i", "04", -1)
			column.Format = strings.Replace(column.Format, "s", "05", -1)
			for _, value := range result {
				date, _ := value[column.Field].(time.Time)
				value[column.Field] = date.Format(column.Format)
			}
		}
	}

	return result
}

// GetModelWheres 获取 model 默认搜索条件
//
//	@receiver s
//	@param c
//	@param model
//	@return string
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

// BuildTree 递归函数，将map数据处理成树形结构
//
//	@receiver s
//	@param items
//	@param parentID
//	@param table
//	@return []map
func (s *sHandleService) BuildTree(items []map[string]interface{}, parentID int64, table tables.BaseTable) []map[string]interface{} {
	node := []map[string]interface{}{}
	// 遍历map，找到所有父ID匹配的项目
	for i := range items {
		if items[i][table.Action.Bind.Recursion.ParentID].(int64) == parentID {
			children := s.BuildTree(items, int64(items[i][table.Action.Bind.Recursion.ChildID].(uint64)), table)
			if len(children) > 0 {
				items[i]["children"] = children
			}
			node = append(node, items[i])
		}
	}
	return node
}
