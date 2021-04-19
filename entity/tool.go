package entity

import (
	"boot/db"
	"boot/protos"
	"gorm.io/gorm"
)

// Tool 刀具信息
type Tool struct {
	gorm.Model
	protos.Tool
}

func init() {
	db.DB.AutoMigrate(&Tool{})
}

func (t *Tool) FindById() *Tool {
	db.DB.First(&t, t.ID)
	return t
}
