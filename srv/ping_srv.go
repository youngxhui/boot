package srv

import (
	pb "boot/protos"
	"context"
	"github.com/youngxhui/power/log"
)

type PingService struct {
}

func (p PingService) Ping(ctx context.Context, request *pb.PingRequest) (*pb.PongReply, error) {
	log.Debug("ping ")
	return &pb.PongReply{Message: "Pong "}, nil
}
