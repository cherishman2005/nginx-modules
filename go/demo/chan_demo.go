package main

import(
	"fmt"
    "time"
)


var (
	GoroutineHandlers int = 5 //多少个协程处理
	nickQueue chan uint64
)

func worker(id int, ch chan uint64) {
    for {
        select {
        case uid := <- ch:
            fmt.Println("handler id=", id, " uid=", uid)
            time.Sleep(1e3)
        }
    }
}

func init() {
    fmt.Println("init")

    nickQueue = make(chan uint64, 1000)
	for i := 0; i < GoroutineHandlers; i++ {
        go worker(i, nickQueue)
	}
}


func main() {
	var WhitelistUids = []uint64{123, 456, 789, 111, 222, 333, 444, 555, 666, 777, 888, 999, 1000}

	for _, v := range WhitelistUids {
		nickQueue <- v
	}
    
    time.Sleep(1e9)
    
}

