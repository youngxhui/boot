package main

import (
	pb "boot/protos"
	"boot/srv"
	"github.com/youngxhui/power/core"
)

func main() {

	c := core.Config{
		IsNeedRegister: false,
		Register:       nil,
	}

	p := core.Power{
		ServerName: "Core",
		Port:       9091,
	}

	s := p.NewServer(c)

	pb.RegisterPingServer(s, &srv.PingService{})
	pb.RegisterUserSrvServer(s, &srv.UserServer{})
	p.Run()
}
