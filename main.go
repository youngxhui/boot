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
		Port:       8080,
	}

	s := p.NewServer(c)

	pb.RegisterGreeterServer(s, &srv.GreeterService{})
	pb.RegisterUserSrvServer(s, &srv.UserServer{})
	p.Run()
}
