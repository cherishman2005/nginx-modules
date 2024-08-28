package main

import (
	"fmt"
)

type Server struct {

}

func (s *Server) f(i int32) {
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

	var i int32 = 11
	sub.Server.f(i)
}
// 子类会使用自己的方法，如果要用父类的方法，就要使用显示父类去调用。
