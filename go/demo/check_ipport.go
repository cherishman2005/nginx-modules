package main

import (
	"log"
	"net"
	"fmt"
	"strings"
)

// func SplitHostPort(hostport string) (host, port string, err error)
func checkIpport(ipport string) bool {
	host, port, err := net.SplitHostPort(ipport)
	if err == nil {
		log.Printf("%s:%s", host, port)
	}

	return err == nil

}

func main() {
	ipstr := "127.0.0.1:4012,127.0.0.1:4013,127.0.0.1:4011,196.168.0.1:4011"
	ips := strings.Split(ipstr, ",")
	if ips == nil || len(ips) == 0 {
		panic("redis config nil")
	}
	for _, v := range ips {
		if !checkIpport(v) {
			panic(fmt.Errorf("SplitHostPort fail with address(%s)", v))
		}
	}
}
