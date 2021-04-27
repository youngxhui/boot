package srv

import (
	"boot/db"
	"boot/protos"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PeopleService struct {
}

// ListPeoples 获取所有的用户
func (p PeopleService) ListPeoples(ctx context.Context, in *protos.ListPeoplesRequest) (*protos.ListPeoplesResponse, error) {
	list, err := db.FindUserList(ctx, int(in.Page), int(in.Size))
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}
	var peoples []*protos.People
	for _, user := range list {
		p := protos.People{
			Id:       int32(user.ID),
			Name:     user.Username,
			Password: user.Password,
			Role:     nil,
		}
		peoples = append(peoples, &p)
	}
	response := protos.ListPeoplesResponse{
		Peoples:       peoples,
	}

	return &response, nil
}

func (p PeopleService) GetPeople(ctx context.Context, in *protos.GetPeopleRequest) (*protos.People, error) {
	return nil, nil
}

// CreatePeople 添加用户
func (p PeopleService) CreatePeople(ctx context.Context, in *protos.CreatePeopleRequest) (*protos.People, error) {
	user, err := db.CreateUser(ctx, in.People.Name, in.People.Password)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}
	return &protos.People{
		Name:     user.Username,
		Password: user.Password,
	}, nil
}

func (p PeopleService) UpdatePeople(ctx context.Context, in *protos.UpdatePeopleRequest) (*protos.People, error) {
	return nil, nil
}

func (p PeopleService) DeletePeople(ctx context.Context, in *protos.DeletePeopleRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
