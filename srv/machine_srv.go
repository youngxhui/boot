package srv

import (
	"boot/db"
	"boot/protos"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MachineService struct {
}

// ListMachines 获取所有的机器信息
func (m MachineService) ListMachines(ctx context.Context, in *protos.ListMachinesRequest) (*protos.ListMachinesResponse, error) {
	list, err := db.FindMachineList(ctx, int(in.Page), int(in.Size))
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}
	var machines []*protos.Machine
	for _, machine := range list {
		m := protos.Machine{
			Id:   int32(machine.ID),
			Name: machine.Name,
		}
		machines = append(machines, &m)
	}
	response := new(protos.ListMachinesResponse)
	response.Machines = machines
	return response, nil
}

// GetMachine 通过id获取设备
func (m MachineService) GetMachine(ctx context.Context, in *protos.GetMachineRequest) (*protos.Machine, error) {
	machine, err := db.FindMachineById(ctx, int(in.Id))
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	return &protos.Machine{
		Id: int32(machine.ID),
		Name: machine.Name,
	}, nil
}
func (m MachineService) CreateMachine(ctx context.Context, in *protos.CreateMachineRequest) (*protos.Machine, error) {
	return nil, nil
}
func (m MachineService) UpdateMachine(ctx context.Context, in *protos.UpdateMachineRequest) (*protos.Machine, error) {
	return nil, nil
}
func (m MachineService) DeleteMachine(ctx context.Context, in *protos.DeleteMachineRequest) (*emptypb.Empty, error) {
	return nil, nil
}
