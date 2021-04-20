package db

import (
	"boot/ent"
	"boot/ent/tool"
	"context"
)

// FindToolById 通过 id 查找 tool
func FindToolById(ctx context.Context, id int) (*ent.Tool, error) {
	t, err := client.Tool.Query().Where(tool.IDEQ(id)).First(ctx)
	return t, err
}

// FindAllTools 获取所有的刀具
func FindAllTools(ctx context.Context, page int, size int) ([]*ent.Tool, error) {
	offSet := (page - 1) * size
	all, err := client.Tool.Query().Limit(size).Offset(offSet).All(ctx)
	return all, err
}
