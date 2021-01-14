package main

import(
	"fmt"
)


var (
	WhitelistUids []uint64
)

func main() {
	WhitelistUids = []uint64{123, 456, 789}
	
	queue := make(chan uint64, 100)
	
	for _, v := range WhitelistUids {
		queue <- v
	}
	
	go SyncYYInfoToBaiDu(queue)
}

/*
func getChanData(c chan uint64) {
	data := <-c
	fmt.Printf("get data :%d\n", data)
}
*/

/*
func SyncYYInfoToBaiDu(c chan uint64) {
	for {
		select {
		case uid := <- c:
			fmt.Println(uid)
			return
		}
	}
}
*/

func SyncYYInfoToBaiDu(c chan uint64) {
	for {
		uid := <- c
		fmt.Println(uid)
	}
}

