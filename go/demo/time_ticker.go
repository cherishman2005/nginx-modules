package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("main start")
	stop := make(chan struct{})

	go loop(stop)

	<- stop
	time.Sleep(10*time.Second)
	fmt.Println("main end")
}

var count int = 0

func loop(stop chan struct{}) {
	t := time.NewTicker(3*time.Second)
	defer func() {
		t.Stop()
		close(stop)
	}()
	for {
		select {
		case <- t.C:
			count++
			if count >= 5 {
				log.Printf("loop count=%d", count)
				return // important: use return and don't use break
			}
			log.Printf("loop incr count=%d", count)
		}
	}
	log.Printf("loop exit count=%d", count)
}


