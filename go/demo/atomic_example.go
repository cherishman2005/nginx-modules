package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
atomic原子操作
*/

var count int64 = 0
//var lock sync.Mutex

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			atomic.AddInt64(&count, 1)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			atomic.AddInt64(&count, 1)
		}
	}()

	wg.Wait()
	fmt.Printf("count=%d\n", count)
}
