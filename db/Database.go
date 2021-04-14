package db

import (
	"github.com/youngxhui/power/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 数据库连接对象
var DB *gorm.DB

func init() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/booster?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		log.Error(err)
		panic(err)
	}

}
