package srv

import (
	pb "boot/protos"
	"context"
)

type GreeterService struct {
}

func (g GreeterService) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello" + request.Name}, nil
}
