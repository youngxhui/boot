package entity

import (
	"boot/db"
	"gorm.io/gorm"
)

const (
	// Health 健康
	Health = 0
	// Waring 即将损坏
	Waring = 1
	// Damage 已经损坏
	Damage = 2
)

// Tool 刀具信息
type Tool struct {
	gorm.Model
	MachineId int
	Status    int
}

//func init() {
//	db.DB.AutoMigrate(&Tool{})
//}

func (t *Tool) FindById() *Tool {
	db.DB.First(&t, t.ID)
	if t.ID == 0 {
		return nil
	}
	return t
}

// FindAllTools 获取所有的Tools
func FindAllTools(page int) []Tool {
	offset := (page - 1) * 10
	var result = db.DB.Limit(10).Offset(offset).Find(&Tool{})
	tools := make([]Tool, result.RowsAffected)
	return tools
}
