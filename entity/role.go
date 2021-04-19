package entity

import (
	"boot/db"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	name string
}

func init() {
	db.DB.AutoMigrate(Role{})
}
