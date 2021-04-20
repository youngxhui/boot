package entity

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	name string
}

//func init() {
//	db.DB.AutoMigrate(Role{})
//}
