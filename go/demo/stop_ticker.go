package main

import (
"fmt"
"time"
)

var hbTicker *time.Ticker
var ch chan int

func main() {
    hbTicker = time.NewTicker(time.Second)
	ch = make(chan int, 10)
	ch <- 10
    go loop()
	
	ch <- 100000
	
	time.Sleep(2*time.Second)
	ch <- -1
    fmt.Println("111")
	time.Sleep(10*time.Second)
	fmt.Println("222")
	hbTicker.Stop()
	hbTicker.Stop()
	time.Sleep(10*time.Second)
	fmt.Println("333")
}

func loop() {
		for {
			select {
			case <-hbTicker.C:
			    fmt.Println("ticker")
			case i := <-ch:
			    fmt.Println("i=", i)
				if i== -1 {
				   fmt.Println("break i=", i)
				   break
				   //return
				}
			}
		}
}
