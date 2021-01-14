package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//取数据
func onlyReadChan(c <-chan int) {
	defer wg.Done()
	data := <-c
	//如果在代码中加入下面的代码，在编译的时候会出现 invalid operation: c <- 12 (send to receive-only type <-chan int)
	//c <- 12
	fmt.Printf("get data: %d\n", data)

}

//写数据
func onlyWriteChan(c chan<- int) {
	defer wg.Done()
	c <- 10
}

func main() {
	c := make(chan int)
	defer close(c)
	wg.Add(2)
	go onlyReadChan(c)
	go onlyWriteChan(c)
	wg.Wait()
	fmt.Printf("end main\n")
}

