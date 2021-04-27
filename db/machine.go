package db

import (
	"boot/ent"
	"boot/ent/machine"
	"context"
)

// FindMachineList 获取所有的machine信息
func FindMachineList(ctx context.Context, page int, size int) ([]*ent.Machine, error) {
	offset := (page - 1) * size
	all, err := client.Machine.Query().Limit(size).Offset(offset).All(ctx)
	return all, err
}

// FindMachineById 通过 id 查找机器
func FindMachineById(ctx context.Context, id int) (*ent.Machine, error) {
	m, err := client.Machine.Query().Where(machine.ID(id)).First(ctx)
	return m, err
}
