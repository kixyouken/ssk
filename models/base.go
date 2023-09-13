package models

import "gorm.io/gorm"

type BaseModel struct {
	gorm.Model
	Table   string                 `gorm:"-"`
	Columns map[string]interface{} `gorm:"-"`
}

func (g *BaseModel) TableName() string {
	return g.Table
}
