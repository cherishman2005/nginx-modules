package main

import (
	"fmt"
	"sync"
)

var count int = 0
var lock sync.Mutex

func main() {
	fmt.Println("hello world");
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			lock.Lock()
			count++
			lock.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			lock.Lock()
			count++
			lock.Unlock()
		}
	}()

	wg.Wait()
	fmt.Printf("count=%d\n", count)
}
