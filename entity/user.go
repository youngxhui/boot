package entity

import (
	"boot/db"
	"github.com/youngxhui/power/log"
	"gorm.io/gorm"
)

// User 数据库 User 表
type User struct {
	gorm.Model
	Username string
	Password string
	RoleId   int32
}

func init() {
	err := db.DB.AutoMigrate(&User{})
	if err != nil {
		panic(err.Error())
	}
}

// Create 对 User 进行添加
func (u User) Create() {
	db.DB.Create(&u)
}

// FindUserByUserNameAndPassword 通过 username 和 password 查询用户
func FindUserByUserNameAndPassword(username, password string) (User, error) {
	log.Debug("database")
	u := User{}
	db.DB.Where("username = ? AND password = ?", username, password).Find(&u)
	return u, nil
}

// FindById 通过 id 查找 user
func (u User)FindById() User {
	db.DB.First(&u)
	return u
}
