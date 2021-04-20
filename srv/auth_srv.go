package srv

import (
	"boot/db"
	"boot/protos"
	"boot/util"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
}

// GetUser 获取当前用户
func (a UserService) GetUser(ctx context.Context, in *protos.GetUserRequest) (*protos.User, error) {
	u, err := db.FindUserById(ctx, int(in.Id))
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}
	return &protos.User{Username: u.Username, Id: int32(u.ID)}, nil
}

// RegisterUser 用户注册
func (a UserService) RegisterUser(ctx context.Context, in *protos.RegisterUserRequest) (*protos.User, error) {
	user, err := db.CreateUser(ctx, in.Username, in.Password)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}
	return &protos.User{
		Username: user.Username,
		Password: user.Password,
	}, nil
}

// LoginUser 用户登录
func (a UserService) LoginUser(ctx context.Context, in *protos.LoginUserRequest) (*protos.LoginUserResponse, error) {
	user, err := db.FindUserByUsernameAndPassword(ctx, in.Username, in.Password)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, "用户名或密码错误")
	}

	token, err := util.GenerateToken(*user)
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, "Token 生失败")
	}

	return &protos.LoginUserResponse{
		Token:  token,
		Header: "Bearer",
	}, status.Error(codes.OK, " 登录成功")

}

// LoginOffUser 用户注销
func (a UserService) LoginOffUser(ctx context.Context, in *protos.DeleteUserRequest) (*protos.LoginOffResponse, error) {

	return nil, nil
}
