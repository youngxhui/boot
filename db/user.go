package db

import (
	"boot/ent"
	"context"
)

func CreateUser(ctx context.Context) (*ent.User, error) {
	a8m, err := client.User. // UserClient.
		Create(). // 用户创建构造器
		SetName("a8m"). // 设置字段的值
		SetAge(18).
		// 设置单个边
		Save(ctx)

	return a8m, err
}
