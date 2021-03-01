package main

import (
	pb "boot/protos"
	"boot/srv"
	"github.com/youngxhui/power/core"
)

func main() {
	p := core.Power{
		ServerName: "Core",
		Port:       8080,
	}

	s := p.NewServer(core.DefaultConfig())

	pb.RegisterGreeterServer(s, &srv.GreeterService{})
	pb.RegisterUserSrvServer(s, &srv.UserServer{})
	p.Run()
}
