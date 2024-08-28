package main

import (
	"fmt"
)

type Server struct {

}

func (s *Server) f(i int) {
	fmt.Println("Server i=", i)
}

type SubServer struct {
	*Server
}

func (s *SubServer) f(i int) {
	fmt.Println("SubServer i=", i)
}

func main() {
	s := &Server{}
	sub := &SubServer{
		Server: s,
	}

	var i int = 11

	sub.f(i)
	sub.Server.f(i)
}
