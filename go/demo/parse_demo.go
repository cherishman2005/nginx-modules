package main

import (
    "fmt"
    "strings"
)

func parseAddrs(addrs string) ([]string){
    Addrs := strings.Split(addrs, ",")

    var cAddrs []string
    for _, addr := range Addrs {
        if len(addr) == 0 {
            continue
        }
        
        cAddrs = append(cAddrs, strings.TrimSpace(addr))
    }
    if len(cAddrs) == 0 {
        cAddrs = []string{"127.0.0.1:6379"}
    }
    return cAddrs
}

func main() {
    addrs := parseAddrs("10.26.135.177:4024, 10.26.135.178:4028 ")
    fmt.Println(addrs)
}