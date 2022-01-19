package main

import (
    "runtime"
    "fmt"
    "time"
)

func getGOMAXPROCS() int {
    return runtime.GOMAXPROCS(0)
}

func cleanLoop() {
    tk := time.NewTicker(time.Second)
    for {
        select {
        case t := <-tk.C:
            now := int32(time.Now().Unix())
            fmt.Println(t.Unix(), now)
        }
    }
}

func main() {
    fmt.Printf("GOMAXPROCS is %d\n", getGOMAXPROCS())
    
    cleanLoop()
}
