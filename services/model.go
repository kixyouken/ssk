package services

import (
	"net/url"
	"ssk/databases"
	"ssk/jsons/models"
	"strconv"
	"strings"

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
func (s *sModelService) GetAll(c *gin.Context, table string, out interface{}, column interface{}, order string) error {
	return db.Table(table).Find(out).Error
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
func (s *sModelService) GetPage(c *gin.Context, out interface{}) error {
	table := FileService.GetTableFile(c)
	order := FileService.GetTableOrders(c, *table)
	model := FileService.GetModelFile(c, table.Action.Bind.Model)
	column := FileService.GetModelColumns(c, *model)
	join := FileService.GetModelJoins(c, *model)

	return db.Table(model.Table.Name).
		Scopes(s.Paginate(c), s.Order(order), s.Joins(join...), s.Search(c), s.Deleted(c, *model)).
		Select(column).Find(out).Error
}

// GetCount 查询数量
//
//	@receiver s
//	@param c
//	@param table 表名
//	@return int64
func (s *sModelService) GetCount(c *gin.Context) int64 {
	table := FileService.GetTableFile(c)
	model := FileService.GetModelFile(c, table.Action.Bind.Model)
	join := FileService.GetModelJoins(c, *model)

	var count int64
	err := db.Table(model.Table.Name).
		Scopes(s.Joins(join...), s.Search(c), s.Deleted(c, *model)).
		Count(&count).Error
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
func (s *sModelService) GetID(c *gin.Context, out interface{}, id int) error {
	form := FileService.GetFormFile(c)
	model := FileService.GetModelFile(c, form.Action.Bind.Model)
	column := FileService.GetModelColumns(c, *model)
	join := FileService.GetModelJoins(c, *model)

	return db.Table(model.Table.Name).
		Scopes(s.Joins(join...), s.Deleted(c, *model)).
		Limit(1).Where(model.Table.Name+".id = ?", id).
		Select(column).Find(out).Error

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

// Deleted 默认删除条件
//
//	@receiver s
//	@param c
//	@return db
//	@return func(db *gorm.DB) *gorm.DB
func (s *sModelService) Deleted(c *gin.Context, model models.BaseModel) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if model.Table.Deleted != nil {
			if model.Table.Deleted.Value != "" {
				db.Where(model.Table.Name+"."+model.Table.Deleted.Field+" = ?", model.Table.Deleted.Value)
			} else {
				db.Where(model.Table.Name + "." + model.Table.Deleted.Field + " IS NULL")
			}
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
		// 获取URL参数
		params, _ := url.QueryUnescape(c.Query("search"))
		// 用 $ 符号拼接搜索各条件
		// search=users.id|eq|100009537$name|like.left|姓名
		if params != "" {
			paramList := strings.Split(params, "$")
			for _, param := range paramList {
				whereList := strings.Split(param, "|")
				switch whereList[1] {
				case "in":
					// search=users.id|in|100009537,100009543
					inList := strings.Split(whereList[2], ",")
					db.Where(whereList[0]+" IN ?", inList)
				case "notin":
					// search=users.id|notin|100009537,100009543
					notinList := strings.Split(whereList[2], ",")
					db.Where(whereList[0]+" NOT IN ?", notinList)
				case "like.left":
					// search=name|like.left|姓名
					db.Where(whereList[0]+" LIKE ?", "%"+whereList[2])
				case "like.right":
					// search=name|like.right|姓名
					db.Where(whereList[0]+" LIKE ?", whereList[2]+"%")
				case "like.all":
					// search=name|like.all|姓名
					db.Where(whereList[0]+" LIKE ?", "%"+whereList[2]+"%")
				case "between.date":
					// search=birthday|between.date|2023-01-01~2023-12-01
					dateList := strings.Split(whereList[2], "~")
					db.Where(whereList[0]+" BETWEEN ? AND ?", dateList[0], dateList[1])
				default:
					db.Where(whereList[0]+s.WhereType(whereList[1]), whereList[2])
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
