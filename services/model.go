package services

import (
	"net/url"
	"ssk/databases"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sModelService struct{}

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
func (s *sModelService) GetPage(c *gin.Context, table string, out interface{}, column interface{}, order string, joins ...string) error {
	return db.Table(table).Scopes(s.Paginate(c), s.Order(order), s.Joins(joins...), s.Search(c)).Select(column).Find(out).Error
}

// GetCount 查询数量
//
//	@receiver s
//	@param c
//	@param table 表名
//	@return int64
func (s *sModelService) GetCount(c *gin.Context, table string, joins ...string) int64 {
	var count int64
	err := db.Table(table).Scopes(s.Joins(joins...), s.Search(c)).Count(&count).Error
	if err != nil {
		return 0
	}

	return count
}

func (s *sModelService) GetID(c *gin.Context, table string, id int, out interface{}, column interface{}) error {
	return db.Table(table).Limit(1).Where("id = ?", id).Select(column).Find(out).Error

}

// Paginate 分页处理
//
//	@receiver s
//	@param c
//	@return db
//	@return func(db *gorm.DB) *gorm.DB
func (s *sModelService) Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page <= 0 {
			page = 1
		}

		limit, _ := strconv.Atoi(c.Query("limit"))
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

func (s *sModelService) Search(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// 获取URL参数
		params, _ := url.QueryUnescape(c.Request.URL.RawQuery)
		if params != "" {
			paramList := strings.Split(params, "&")
			for _, param := range paramList {
				whereList := strings.Split(param, "|")
				switch whereList[1] {
				case "like.left":
					db.Where(whereList[0]+" LIKE ?", "%"+whereList[2])
				case "like.right":
					db.Where(whereList[0]+" LIKE ?", whereList[2]+"%")
				case "like.all":
					db.Where(whereList[0]+" LIKE ?", "%"+whereList[2]+"%")
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
	case "lt":
		where = " < ?"
	case "gt":
		where = " > ?"
	}
	return where
}
