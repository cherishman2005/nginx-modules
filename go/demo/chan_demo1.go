package main

import(
	"fmt"
    "time"
)


var (
	GoroutineHandlers int = 5 //多少个协程处理
	queue chan *StoreUserInfo
)

type StoreUserInfo struct {
	uid uint64
	nick string
	version uint64
}

func worker(id int, ch chan *StoreUserInfo) {
    for {
        select {
        case storeUserInfo := <- ch:
            fmt.Println("worker id=", id, " storeUserInfo=", storeUserInfo.uid, storeUserInfo.nick, storeUserInfo.version)
            time.Sleep(1e3)
        }
    }
}

func init() {
    fmt.Println("init")

    queue = make(chan *StoreUserInfo, 1000)
	for i := 0; i < GoroutineHandlers; i++ {
        go worker(i, queue)
	}
}


func main() {
	var WhitelistUids = []uint64{123, 456, 789, 111, 222, 333, 444, 555, 666, 777, 888, 999, 1000}

	for _, v := range WhitelistUids {
		storeUserInfo := &StoreUserInfo{
		  uid: v,
		  nick: "zhangbiwu",
		  version: 66777,
		}
		queue <- storeUserInfo
	}
    
    time.Sleep(1e9)
    
}

