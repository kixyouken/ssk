package services

import (
	"ssk/databases"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db = databases.InitMysql()

func GetAll(c *gin.Context, table string, out interface{}) error {
	return db.Table(table).Find(out).Error
}

func GetPage(c *gin.Context, table string, out interface{}, column interface{}, order string) error {
	return db.Table(table).Scopes(Paginate(c), Order(order)).Select(column).Find(out).Error
}

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
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

func Order(order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if order == "" {
			order = "id DESC"
		}
		return db.Order(order)
	}
}
