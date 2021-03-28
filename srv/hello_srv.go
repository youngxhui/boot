package srv

import (
	pb "boot/protos"
	"context"
	"github.com/youngxhui/power/log"
)

type GreeterService struct {
}

func (g GreeterService) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Debug("greeter service " + request.Name)
	return &pb.HelloReply{Message: "Hello " + request.Name}, nil
}
