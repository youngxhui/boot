package main

import "github.com/youngxhui/power/core"

func main() {
	p := core.Power{
		ServerName: "Core",
		Port:       8080,
	}
	s := p.NewServer(core.DefaultConfig())

	p.Run()
}
