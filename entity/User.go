package entity

import (
	"boot/db"
	"gorm.io/gorm"
)

// User 数据库 User 表
type User struct {
	gorm.Model
	Username string
	Password string
}

// Create 对 User 进行添加
func (u User) Create() {
	db.DB.Create(&u)
}