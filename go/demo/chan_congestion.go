package main

import (
"fmt"
"time"
)

var hbTicker *time.Ticker
var ch chan int
var start chan bool

func main() {
    //hbTicker = time.NewTicker(time.Second)
	ch = make(chan int, 10)
	start = make(chan bool, 1)
	
	go loop()
	
	for i := 0; i < 20; i++ {
	    ch <- i
		fmt.Println("i=", i, "len=", len(ch))
		if len(ch) >= 10 {
		    start <- true
		}
	}
	
	fmt.Println("end len=", len(ch))

	time.Sleep(10*time.Second)
	fmt.Println("333")
	

	
}

func loop() {
        <- start
		for {
			select {
			//case <-hbTicker.C:
			//    fmt.Println("ticker")
			case i := <-ch:
			    fmt.Println("consume i=", i)
			}
		}
}
