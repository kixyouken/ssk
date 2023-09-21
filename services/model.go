package services

import (
	"ssk/databases"
	"ssk/jsons/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sModelService struct{}

type JoinColumns struct {
	Field string `json:"field"`
}

var ModelService = sModelService{}

var db = databases.InitMysql()

// GetAll 查询所有
//
//	@receiver s
//	@param c
//	@param table 表名
//	@param out
//	@param column 字段
//	@param order 排序
//	@return error
func (s *sModelService) GetAll(c *gin.Context, table string, out interface{}, column interface{}, order string, search interface{}) error {
	return db.Table(table).Where(search).
		Scopes(s.Order(order)).Select(column).
		Find(out).Error
}

// GetPage 分页查询
//
//	@receiver s
//	@param c
//	@param table 表名
//	@param out
//	@param column 字段
//	@param order 排序
//	@return error
func (s *sModelService) GetPage(c *gin.Context, table string, out interface{}, column interface{}, order string, join []string, search interface{}) error {
	return db.Table(table).Where(search).
		Scopes(s.Paginate(c), s.Order(order), s.Joins(join...), s.Search(c)).
		Select(column).Find(out).Error
}

// GetCount 查询数量
//
//	@receiver s
//	@param c
//	@param table 表名
//	@return int64
func (s *sModelService) GetCount(c *gin.Context, table string, join []string, search interface{}) int64 {
	var count int64
	err := db.Table(table).Where(search).
		Scopes(s.Joins(join...), s.Search(c)).
		Count(&count).Error
	if err != nil {
		return 0
	}

	return count
}

// GetWithsCount 关联数量
//
//	@receiver s
//	@param c
//	@param table
//	@param search
//	@return int64
func (s *sModelService) GetWithsCount(c *gin.Context, table string, search interface{}) int64 {
	var count int64
	err := db.Table(table).Where(search).Count(&count).Error
	if err != nil {
		return 0
	}

	return count
}

// GetID 根据ID查询
//
//	@receiver s
//	@param c
//	@param table 表名
//	@param id
//	@param out
//	@param column 字段
//	@return error
func (s *sModelService) GetID(c *gin.Context, id int, table string, out interface{}, column interface{}, join []string) error {
	return db.Table(table).Scopes(s.Joins(join...)).
		Limit(1).Where(table+".id = ?", id).
		Select(column).Find(out).Error

}

// GetDistinctCount 获取去重统计数量
//
//	@receiver s
//	@param c
//	@param table
//	@param fields
//	@return int64
func (s *sModelService) GetDistinctCount(c *gin.Context, table string, fields string) int64 {
	var count int64
	err := db.Table(table).Distinct(fields).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

// GetDistinct 获取去重数据
//
//	@receiver s
//	@param c
//	@param table
//	@param out
//	@param fields
//	@return error
func (s *sModelService) GetDistinct(c *gin.Context, table string, out interface{}, fields string) error {
	return db.Raw("SELECT DISTINCT ( " + fields + " ) FROM " + table).Scan(out).Error
}

// GetSql 原生 sql 查询
//
//	@receiver s
//	@param c
//	@param sql
//	@param out
//	@return error
func (s *sModelService) GetSql(c *gin.Context, sql string, out interface{}) error {
	return db.Raw(sql).Scan(out).Error
}

// SetCreate 添加数据
//
//	@receiver s
//	@param c
//	@param table
//	@param data
//	@return error
func (s *sModelService) SetCreate(c *gin.Context, table string, param map[string]interface{}) error {
	return db.Table(table).Create(param).Error
}

// SetUpdate 更新数据
//
//	@receiver s
//	@param c
//	@param table
//	@param id
//	@param updates
//	@return error
func (s *sModelService) SetUpdate(c *gin.Context, table string, id int, updates interface{}) error {
	return db.Table(table).Where("id = ?", id).Updates(updates).Error
}

// SetDelete 删除数据
//
//	@receiver s
//	@param c
//	@param table
//	@param id
//	@param model
//	@return error
func (s *sModelService) SetDelete(c *gin.Context, table string, id int, model models.BaseModel) error {
	if model.Table.Deleted != nil {
		switch model.Table.Deleted.Value {
		case "time":
			return db.Table(table).Where("id = ?", id).Update(model.Table.Deleted.Field, time.Now()).Error
		default:
			return db.Table(table).Where("id = ?", id).Update(model.Table.Deleted.Field, model.Table.Deleted.Value).Error
		}
	} else {
		return db.Table(table).Where("id = ?", id).Update("deleted_at", time.Now()).Error
	}
}

// Paginate 分页处理
//
//	@receiver s
//	@param c
//	@return db
//	@return func(db *gorm.DB) *gorm.DB
func (s *sModelService) Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		table := FileService.GetTableFile(c)
		page := table.Action.Page
		if table.Action.Page <= 0 {
			page = 1
		}

		urlPage, _ := strconv.Atoi(c.Query("page"))
		if urlPage > 0 {
			page = urlPage
		}

		limit := table.Action.Limit
		if limit > int(table.Action.Count) {
			limit = int(table.Action.Count)
		}
		urlLimit, _ := strconv.Atoi(c.Query("limit"))
		if urlLimit > 0 {
			limit = urlLimit
		}
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

// Order 排序处理
//
//	@receiver s
//	@param order
//	@return db
//	@return func(db *gorm.DB) *gorm.DB
func (s *sModelService) Order(order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

// Joins 模型关联处理
//
//	@receiver s
//	@param joins
//	@return db
//	@return func(db *gorm.DB) *gorm.DB
func (s *sModelService) Joins(joins ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, v := range joins {
			db.Joins(v)
		}
		return db
	}
}

// Columns 获取关联表字段信息
//
//	@receiver s
//	@param c
//	@param model
//	@return []string
func (s *sModelService) JoinColumns(c *gin.Context, model models.BaseModel) []string {
	columns := []string{}
	for _, value := range model.Table.Joins {
		if value.Columns == nil || len(value.Columns) == 0 {
			joinColumns := []JoinColumns{}
			db.Raw("SHOW COLUMNS FROM `" + value.Name + "`").Scan(&joinColumns)
			for _, v := range joinColumns {
				columns = append(columns, value.Name+"."+v.Field+" AS SSK_"+value.Name+"_"+v.Field)
			}
		} else {
			for _, v := range value.Columns {
				if strings.Contains(v.Field, "as") || strings.Contains(v.Field, "AS") {
					columns = append(columns, value.Name+"."+v.Field)
				} else {
					columns = append(columns, value.Name+"."+v.Field+" AS "+value.Name+"_"+v.Field)
				}
			}
		}
	}

	return columns
}

// Search 搜索条件处理
//
//	@receiver s
//	@param c
//	@return db
//	@return func(db *gorm.DB) *gorm.DB
func (s *sModelService) Search(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		table := FileService.GetTableFile(c)
		model := FileService.GetModelFile(c, table.Action.Bind.Model)
		// ?name=like.left.test&master_university.id=eq.2
		search := c.Request.URL.Query()
		for k, v := range search {
			lastDotIndex := strings.LastIndex(v[0], ".")
			if lastDotIndex != -1 {
				where := v[0][:lastDotIndex]
				value := v[0][lastDotIndex+1:]
				if !strings.Contains(k, ".") {
					k = model.Table.Name + "." + k
				}
				switch where {
				case "in":
					inList := strings.Split(value, ",")
					db.Where(k+" IN ?", inList)
				case "notin":
					notinList := strings.Split(value, ",")
					db.Where(k+" NOT IN ?", notinList)
				case "like.left":
					db.Where(k+" LIKE ?", "%"+value)
				case "like.right":
					db.Where(k+" LIKE ?", value+"%")
				case "like":
					db.Where(k+" LIKE ?", "%"+value+"%")
				case "between.date":
					dateList := strings.Split(value, "~")
					db.Where(k+" BETWEEN ? AND ?", dateList[0], dateList[1])
				default:
					db.Where(k+s.WhereType(where), value)
				}
			}
		}
		return db
	}
}

func (s *sModelService) WhereType(where string) string {
	switch where {
	case "eq":
		where = " = ?"
	case "neq":
		where = " <> ?"
	case "lt":
		where = " < ?"
	case "elt":
		where = " <= ?"
	case "gt":
		where = " > ?"
	case "egt":
		where = " >= ?"
	}
	return where
}
