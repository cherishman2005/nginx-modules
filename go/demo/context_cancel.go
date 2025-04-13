package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, "worker1")
	go worker(ctx, "worker2")

	time.Sleep(1 * time.Second)
	cancel() // 取消所有工作

	time.Sleep(1 * time.Second)
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "stopped")
			return
		default:
			fmt.Println(name, "working")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
