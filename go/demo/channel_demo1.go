package main

import (
	"fmt"
)

func fibonacci(c chan int, quit chan bool) {
	for {
		select {
		case i := <- c:
			fmt.Println(i)
			// process
		case <-quit:
			return
		}
	}
}

func main() {
	n := 10
	queue := make(chan int)
	quit := make(chan bool, 1)
	

	go func() {
		for i := 0; i < n; i++ {
			queue <- i
		}
		quit <-false
	}()

	fibonacci(queue, quit)
}