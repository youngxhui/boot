package db

import (
	"boot/ent"
	"boot/ent/user"
	"context"
)

// CreateUser 创建 user
func CreateUser(ctx context.Context, username, password string) (*ent.User, error) {
	u, err := client.User. // UserClient.
		Create(). // 用户创建构造器
		SetUsername(username). // 设置字段的值
		SetPassword(password).
		// 设置单个边
		Save(ctx)

	return u, err
}

// FindUserByUsernameAndPassword 通过用户名和密码查询用户
func FindUserByUsernameAndPassword(ctx context.Context, username, password string) (*ent.User, error) {
	first, err := client.User.Query().Where(user.UsernameEQ(username), user.PasswordEQ(password)).First(ctx)
	return first, err
}

// FindUserById 通过id查找用户
func FindUserById(ctx context.Context, id int) (*ent.User, error) {
	u, err := client.User.Query().Where(user.IDEQ(id)).First(ctx)
	return u, err
}
