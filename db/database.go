package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// DB 数据库连接对象
var DB *gorm.DB

func init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	// "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//dsn := "root:1234@tcp(127.0.0.1:3306)/booster?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "host=localhost user=kong password=kong dbname=boot port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true, Logger: newLogger}); err != nil {

		panic(err)
	}

}
