package srv

import (
	"boot/protos"
	"context"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/youngxhui/power/log"
)
// UserServer 用户服务
type UserServer struct {
}

// Login 登录服务
func (u UserServer) Login(ctx context.Context, info *protos.UserInfo) (*protos.Result, error) {
	if info.Username == "123" && info.Password == "1234" {
		return &protos.Result{
			Code:    0,
			Message: "登录成功",
			Data:    nil,
		}, nil
	}


	token := protos.Success{
		Token: "Java",
	}
	any, _ := anypb.New(&token)
	log.Info("any:", any)
	newToken := protos.Success{}

	log.Info("new Token", newToken.Token)
	return nil, status.Error(16, "用户名或密码错误")
}
