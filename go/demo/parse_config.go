package main

import (
    "fmt"
    "io"
    "os"
    //"flag"
    //"sync"
    //"sync/atomic"
    //"context"
    //"time"
    "strings"
    "bufio"
    "strconv"
)

var CallIf = make(map[string]int)

func readLine(pathname string) {
    fi, err := os.Open(pathname)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }
    defer fi.Close()

    br := bufio.NewReader(fi)
    for {
        a, _, c := br.ReadLine()
        if c == io.EOF {
            break
        }
        fmt.Println(string(a))
        
        kv := strings.Split(string(a), ":")
        if len(kv) != 2 {
            continue
        }
        
        k := strings.TrimSpace(kv[0])
        if len(k) == 0 {
            continue
        }
        
        v := strings.TrimSpace(kv[1])
        //n, err := strconv.ParseInt(v, 10, 64)
        n, err := strconv.Atoi(v)
        if err != nil {
            continue
        }
        CallIf[k] = n
    }
    
    return
}

func main() {
    readLine("./call_interface.txt")
    fmt.Println(CallIf)
}

/*
./call_interface.txt示例

aaa:1
bbb:1
ccc: 111
ddd:1

*/