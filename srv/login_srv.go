package srv

import (
	"boot/protos"
	"context"
)

type UserServer struct {
}

func (u UserServer) Login(ctx context.Context, info *protos.UserInfo) (*protos.Result, error) {

	return &protos.Result{}, nil
}
