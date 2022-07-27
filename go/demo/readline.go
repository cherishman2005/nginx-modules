package main

import (
    "fmt"
    "os"
    "io"
    "bufio"
)

func ReadIplist(pathname string) (map[string]bool) {
    iplist := make(map[string]bool)

    fi, err := os.Open(pathname)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return iplist
    }
    defer fi.Close()

    br := bufio.NewReader(fi)
    for {
        a, _, c := br.ReadLine()
        if c == io.EOF {
            break
        }
        //fmt.Println(string(a))
        if len(a) > 0 {
            iplist[string(a)] = true
        }
    }
    
    return iplist
}


func main() {
    pathname := "./ip_whitelist.txt"
    iplist := ReadIplist(pathname)
    fmt.Println(iplist)
    for ip,_ := range iplist {
        fmt.Println(ip)
    }
}