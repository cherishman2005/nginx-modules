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
    timer := time.NewTimer(time.Second)
    select {
    case t := <-timer.C:
        //shard.onCleanTimer(t)
        now := int32(time.Now().Unix())
        fmt.Printf("t:%v now:%d\n", t, now)
    }
    timer.Stop()
}

func main() {
    fmt.Printf("GOMAXPROCS is %d\n", getGOMAXPROCS())
    
    cleanLoop()
    
    timer := time.NewTimer(3 * time.Second) 
    for {
        timer.Reset(3 * time.Second) // 这里复用timer
        select {
        case <-timer.C:
            fmt.Println("Execute every 3 seconds")
        }
    }
}