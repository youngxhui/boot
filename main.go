package main

import (
	_ "boot/db"
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
	pb.RegisterUserServiceServer(s, &srv.UserService{})
	p.Run()
}
